package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int    `json:"id"`
	EMAIL string `json:"email"`
	NOME  string `json:"name"`
}

type Endereco struct {
	ID     int    `json:"id"`
	CIDADE string `json:"cidade"`
	BAIRRO string `json:"bairro"`
	RUA    string `json:"rua"`
	IDUSER int    `json:"id_user"`
}

type Livros struct {
	NOME   string `json:"name"`
	AUTOR  string `json:"autor"`
	GENERO string `json:"genero"`
	QTD    int    `json:"quantidade"`
}

type Estante struct {
	ID      int `json:"id"`
	IDUSER  int `json:"id_user"`
	IDLIVRO int `json:"id_livro"`
}

var db, err = sql.Open("mysql", "leandro:123456@/biblioteca")

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, er := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && er == nil {
		fmt.Print("token verificado")
	} else {
		result := gin.H{
			"message": "nao autorizado",
			"error":   er.Error(),
		}
		c.JSON(200, result)
		c.Abort()
	}
}

func isLogin(c *gin.Context) {
	reqEmail := c.PostForm("email")
	reqPassword := c.PostForm("password")

	senha := encriptar(reqPassword)

	var email string
	var password string
	db.QueryRow("SELECT email,password FROM users WHERE email = ? AND password = ?", reqEmail, senha).Scan(&email, &password)

	if len(email) > 0 && len(senha) > 0 {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": email,
			"password": password,
		})

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			log.Fatal(err)

		}
		c.JSON(200, tokenString)
	} else {
		c.JSON(401, "Email ou Senha Incorretos")
	}

}

func getLivros(c *gin.Context) {
	id := c.Param("id")
	rows, er := db.Query(`
	SELECT ll.nome, LL.AUTOR, g.nome, LL.QTD
	FROM users u 
	INNER JOIN ligacao1 l
	ON l.id_user = u.id
	JOIN livros ll
	ON l.id_livros = ll.id
	JOIN generos g
	ON l.id_genero = g.id
	WHERE l.id_user = ?
	`, id)

	if er != nil {
		log.Fatal(er)

	}

	var livros []Livros
	for rows.Next() {
		l := Livros{}
		e := rows.Scan(&l.NOME, &l.AUTOR, &l.GENERO, &l.QTD)
		if e != nil {
			log.Fatal(e)
		}

		livros = append(livros, l)
	}

	c.JSON(200, livros)

}

func deleteLigation(id int64) {
	stmtLigacao, _ := db.Prepare("DELETE FROM ligacao1 WHERE id_user = ?")
	stmtLigacao.Exec(id)
}

func deleteLivroById(id int64) {
	stmt, _ := db.Prepare("DELETE FROM livros WHERE id_user = ?")
	stmt.Exec(id)
}

func deleteEstante(id int64) {
	stmt, _ := db.Prepare("DELETE FROM estante WHERE id_user = ?")
	stmt.Exec(id)
}

func deleteEndereco(id int64) {
	stmt, _ := db.Prepare("DELETE FROM endereco WHERE id_user = ?")
	stmt.Exec(id)
}

func getId(email string) int {
	var id int
	db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&id)
	return id
}

func encriptar(value string) string {
	passByte := md5.New()
	passByte.Write([]byte(value))
	return hex.EncodeToString(passByte.Sum(nil))
}

func getUser(c *gin.Context) {
	email := c.Param("email")
	var user User

	db.QueryRow("SELECT id,nome,email FROM users WHERE email = ?", email).Scan(&user.ID, &user.NOME, &user.EMAIL)

	c.JSON(200, user)
}

func cadastroUser(c *gin.Context) {
	Nome := c.PostForm("nome")
	Email := c.PostForm("email")
	Password := c.PostForm("password")
	Cidade := c.PostForm("cidade")
	Bairro := c.PostForm("bairro")
	Rua := c.PostForm("rua")

	senha := encriptar(Password)

	var n string
	db.QueryRow("SELECT email FROM users WHERE email = ?", Email).Scan(&n)

	if len(n) > 0 {
		c.JSON(409, "Ja existe este Usuario")
		return
	}

	stmt, err := db.Prepare("INSERT INTO users(NOME, EMAIL, PASSWORD) VALUES(?,?,?)")

	if err != nil {
		log.Fatal(err)
		c.JSON(500, err.Error())
	}

	stmt.Exec(Nome, Email, senha)

	id := getId(Email)

	stmtEndereco, errE := db.Prepare("INSERT INTO endereco(cidade, bairro, rua, id_user) VALUES(?,?,?,?)")

	if errE != nil {
		log.Fatal(errE)
		c.JSON(500, errE.Error())
		return
	}

	stmtEndereco.Exec(Cidade, Bairro, Rua, id)

	var idEnd int
	db.QueryRow("SELECT id FROM endereco WHERE id_user = ?", id).Scan(&idEnd)

	stmtLigacao, _ := db.Prepare("INSERT INTO ligacao1(id_user,id_endereco) VALUES(?,?)")

	stmtLigacao.Exec(id, idEnd)

	c.JSON(200, "Cadastrado com sucesso")

}

func cadastrarLivro(c *gin.Context) {
	id := c.Param("id")
	Nome := c.PostForm("nome")
	Autor := c.PostForm("autor")
	GeneroS := c.PostForm("genero")
	qtd := c.PostForm("qtd")
	quantidade, _ := strconv.ParseInt(qtd, 10, 11)
	Genero, _ := strconv.ParseInt(GeneroS, 10, 11)

	var livro string
	db.QueryRow("SELECT nome FROM livros WHERE id_user = ? AND nome = ?", id, Nome).Scan(&livro)

	if len(livro) > 0 {
		c.JSON(401, "Ja existe "+Nome+" em sua estante")
		return
	}

	stmt, err := db.Prepare("INSERT INTO livros(id_user, nome, autor, qtd) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	stmt.Exec(id, Nome, Autor, quantidade)

	var idLivro int

	db.QueryRow("SELECT id FROM livros WHERE id_user = ? AND nome = ?", id, Nome).Scan(&idLivro)

	stmtEstante, _ := db.Prepare("INSERT INTO estante(id_user,id_livro) VALUES(?,?)")
	stmtEstante.Exec(id, idLivro)

	var idEst int
	db.QueryRow("SELECT id FROM estante WHERE id_user = ?", id).Scan(&idEst)

	stmtLig, _ := db.Prepare("INSERT INTO ligacao1(id_user, id_estante, id_livros, id_genero) VALUES(?,?,?,?)")
	stmtLig.Exec(id, idEst, idLivro, Genero)

	c.JSON(200, "Livro "+Nome+" Adicionado com sucesso")
}

func deleteLivro(c *gin.Context) {
	idS := c.Param("id")
	Nome := c.Param("livro")

	id, _ := strconv.ParseInt(idS, 10, 8)
	var livro int
	db.QueryRow("SELECT id from livros where id_user = ? AND nome =?", id, Nome).Scan(&livro)

	if livro == 0 {
		c.JSON(404, "Nao existe esse livro")
		return
	}

	stmte, e := db.Prepare("DELETE FROM livros WHERE  livros.id_user = ? AND nome = ?")

	if e != nil {
		log.Fatal(e)
		fmt.Print(e)
		return
	}

	stmte.Exec(id, Nome)

	stmL, er := db.Prepare("UPDATE ligacao1 SET id_livros = null WHERE id_user = ? AND id_livros = ?")

	if er != nil {
		log.Fatal(er)
		return
	}

	stmL.Exec(id, livro)

	stmE, er := db.Prepare("DELETE FROM estante WHERE id_user = ?, id_livro = ?")

	if er != nil {
		log.Fatal(er)
		return
	}

	stmE.Exec(id, livro)

	c.JSON(200, "Excluido com sucesso")
}

func deletarUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	senha := encriptar(password)

	var id int64
	db.QueryRow("SELECT id FROM users WHERE email = ? AND password = ?", email, senha).Scan(&id)

	if id == 0 {
		c.JSON(404, "Usuario nao encontrado")
		return
	}

	deleteLigation(id)
	deleteLivroById(id)
	deleteEstante(id)
	deleteEndereco(id)

	stmt, er := db.Prepare("DELETE FROM users WHERE id = ?")

	if er != nil {
		log.Fatal(er)
		return
	}

	stmt.Exec(id)

	c.JSON(200, "Usuario Excluido com sucesso!")
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/cadastro/", cadastroUser)
	router.POST("/cadastroLivro/:id", auth, cadastrarLivro)
	router.POST("/deleteUser/", auth, deletarUser)
	router.POST("/isLogin/", isLogin)
	router.GET("/getLivros/:id", auth, getLivros)
	router.GET("/getUser/:email", auth, getUser)
	router.DELETE("/deleteLivro/:id/:livro", auth, deleteLivro)
	router.Run(":9000")
}
