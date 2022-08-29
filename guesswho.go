//This is a basic game - The user must "guess a name" of a character and they can either ask for a tip or guess right away. The player starts with 6 points. Every tip costs 1 point, every mistake costs 2 points.
//Soon I will recreate it, using JavaScript and HTML, for a web - and more interactive - version.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type NomesChar struct {
	Nome  string
	Nome2 string
	Nome3 string
}

type DicasChar struct {
	Dica1 string
	Dica2 string
	Dica3 string
	Dica4 string
}

type Character struct {
	NomesChar
	DicasChar
}

var last string = " Esta é a sua última dica."
var dc string = " Sou da DC. "
var marvel string = " Sou da Marvel. "
var mito string = " Sou um herói da mitologia. "

func SeverChar() Character {

	CharA := Character{
		NomesChar{"BATMAN", "THE BATMAN", "HOMEM MORCEGO"},
		DicasChar{dc, " Também sou investigador. ", " Sou considerado uma das pessoas mais inteligentes da terra. ", " Sou órfão. " + last},
	}
	CharB := Character{
		NomesChar{"SUPERMAN", "SUPER MAN", "SUPER HOMEM"},
		DicasChar{dc, " Não sou da terra. ", " Meu principal uniforme é azul. ", " Meu nome também é Kalel. " + last},
	}
	CharC := Character{
		NomesChar{"SHAZAM", "CHAZAM", "XAZAM"},
		DicasChar{dc, " Sou uma criança ", " Me transformo em adulto. ", " Meu uniforme leva tem um raio amarelo no peito. " + last},
	}
	CharD := Character{
		NomesChar{"MULHER MARAVILHA", "WONDER WOMAN", "MULHERMARAVILHA"},
		DicasChar{dc, " Nasci em uma ilha isolada. ", " Uma das minhas armas é um chicote. ", " Sou uma amazona. " + last},
	}
	CharE := Character{
		NomesChar{"FLASH", "FLESH", "THE FLASH"},
		DicasChar{dc, " Meu uniforme é vermelho. ", " Meu nome verdadeiro é Barry Allen. ", " Sou muito veloz. " + last},
	}
	CharF := Character{
		NomesChar{"LEX LUTHOR", "LEX LUTOR", "LAX LUTOR"},
		DicasChar{dc, " Sou careca. ", " Por vezes uso uma armadura robótica. ", " Sou o principal inimigo do Super Homem. " + last},
	}
	CharG := Character{
		NomesChar{"CORINGA", "JOKER", "CURINGA"},
		DicasChar{dc, " HA HA HA HA! ", " Gosto de fazer as pessoas sorrirem. ", " O Batman não vai com a minha cara. " + last},
	}
	CharH := Character{
		NomesChar{"AQUAMAN", "AQUA MAN", "AQUAMEN"},
		DicasChar{dc, " GLUB. ", " Meu reino é escondido. ", " Falo com os animais aquáticos. " + last},
	}
	CharI := Character{
		NomesChar{"HULK", "THE HULK", "O HULK"},
		DicasChar{marvel, " ESMAGA! ", " Sou um dos personagens mais fortes. ", " Sou verde. " + last},
	}
	CharJ := Character{
		NomesChar{"CAPITÃO AMÉRICA", "CAPITAO AMERICA", "CAPITÃO AMERICA"},
		DicasChar{marvel, " Lutei na Segunda Guerra Mundial. ", " Fui congelado por muito tempo. ", " Uso um escudo. " + last},
	}
	CharK := Character{
		NomesChar{"HOMEM-ARANHA", "SPIDERMAN", "HOMEM ARANHA"},
		DicasChar{marvel, " Meu tio não está Ben. ", " Sou meio nerd. ", " Escalo paredes. " + last},
	}
	CharL := Character{
		NomesChar{"IRON MAN", "HOMEM DE FERRO", "IRONMAN"},
		DicasChar{marvel, " Sou muito rico. ", " Meus projetos ganham o nome de 'Mark'. ", " Sou de ferro. " + last},
	}
	CharM := Character{
		NomesChar{"VIÚVA NEGRA", "VIUVA NEGRA", "VIÚVA NEGRA"},
		DicasChar{marvel, " Fui treinada na Rússia. ", " Bart é meu amigo. ", " Tenho o nome de uma aranha. " + last},
	}
	CharN := Character{
		NomesChar{"GAVIÃO ARQUEIRO", "GAVIAO ARQUEIRO", "HAWKEYE"},
		DicasChar{marvel, " Também me tornei um ronin. ", " Muitos me acham inútil. ", " Uso arco e flecha. " + last},
	}
	CharO := Character{
		NomesChar{"THANOS", "TANOS", "TANUS"},
		DicasChar{marvel, " Reuni as joias do infinito ", " Dizimei metade da população do universo. ", " Você deveria ter acertado a cabeça. " + last},
	}
	CharP := Character{
		NomesChar{"THOR", "TOR", "TÓR"},
		DicasChar{marvel, " Não sou da terra. ", " Ele é um amigo! Do trabalho! ", " Meu irmão é meio problemático. " + last},
	}
	CharQ := Character{
		NomesChar{"LOKI", "LÓKI", "LOQUI"},
		DicasChar{marvel, " Um ser superior como eu não deveria ficar dando dicas. ", " Minha verdadeira família é azul. ", " Sou o mestre das ilusões. " + last},
	}
	CharR := Character{
		NomesChar{"ULTRON", "ULTROM", "ÚLTRON"},
		DicasChar{marvel, " EU NÃO SOU O PINOCCHIO! ", " Meu pai é o Tony Stark. ", " Sou feito de um metal muito resistente. " + last},
	}
	CharS := Character{
		NomesChar{"DOUTOR ESTRANHO", "DR ESTRANHO", "DR. ESTRANHO"},
		DicasChar{marvel, " A magia me guia. ", " Sempre abro portais. ", " Ando sempre com o Manto da Levitação. " + last},
	}
	CharT := Character{
		NomesChar{"VENOM", "VENON", "VENNOM"},
		DicasChar{marvel, " Nós somos um! ", " Fazemos simbiose. ", " Odeio o homem-aranha. " + last},
	}
	CharU := Character{
		NomesChar{"PANTERA NEGRA", "BLACK PANTHER", "O PANTERA NEGRA"},
		DicasChar{marvel, " Sou o mais rico de meu universo. ", " Meu país é oculto. ", " Meu poder é ancestral. " + last},
	}
	CharV := Character{
		NomesChar{"HÉRCULES", "HERCULES", "HERCULIS"},
		DicasChar{mito, " Cumpri 12 trabalhos. ", " Matei minha própria família. ", " Tenho a força de 300 homens. " + last},
	}
	CharW := Character{
		NomesChar{"AQUILES", "AQUILIS", "AQUILES"},
		DicasChar{mito, " Sou o mais belo dos heróis. ", " Sou quase invencível. ", " Minha fraqueza é meu calcanhar. " + last},
	}
	CharX := Character{
		NomesChar{"TESEU", "TESEO", "THESEU"},
		DicasChar{mito, " Enfrentei um minotauro. ", " Tornei-me rei. ", " Sou tão forte quanto Hércules. " + last},
	}
	CharY := Character{
		NomesChar{"PERSEU", "PERSEL", "PERSEO"},
		DicasChar{mito, " Derrotei uma poderosa górgona. ", " Já usei meu escudo para não virar pedra. ", " Sou filho de Zeus. " + last},
	}
	CharZ := Character{
		NomesChar{"ORFEU", "ORFEL", "ORFEO"},
		DicasChar{mito, " Também sou músico. ", " Sou um Argonauta. ", " Já fui ao submundo para resgatar a minha eposa. " + last},
	}

	rand.Seed(time.Now().Unix())
	RandomChar := []Character{

		CharA, CharB, CharC, CharD, CharE, CharF,
		CharG, CharH, CharI, CharJ, CharK, CharL,
		CharM, CharN, CharO, CharP, CharQ, CharR,
		CharS, CharT, CharU, CharV, CharW, CharX,
		CharY, CharZ,
	}
	return RandomChar[rand.Intn(len(RandomChar))]
}

func (c *Character) SetChar(Character) (Selecionado Character) {

	c.Nome = SeverChar().Nome
	c.Nome2 = SeverChar().Nome2
	c.Nome3 = SeverChar().Nome3
	c.Dica1 = SeverChar().Dica1
	c.Dica2 = SeverChar().Dica2
	c.Dica3 = SeverChar().Dica3
	c.Dica4 = SeverChar().Dica4

	return
}

var respErr string = " Errrouuuu! "
var pont string = " Sua pontuação é: "
var tip string = " A dica é:"
var Cong string = " Parabéns! Você acertou!"
var x, a, t, e int8 = 6, 0, 1, 2

func string1(first string, Selecionado Character) (resposta1 string, ponto1 int8) {

	if first == Selecionado.Nome || first == Selecionado.Nome2 || first == Selecionado.Nome3 {
		return Cong + pont, x - a
	} else if first == "TIP" {
		return tip + Selecionado.Dica1 + pont, x - t
	} else {
		return respErr + pont, x - e
	}
}

func string2(second string, Selecionado Character, ponto1 int8) (resposta2 string, ponto2 int8) {
	x = ponto1
	if second == Selecionado.Nome || second == Selecionado.Nome2 || second == Selecionado.Nome3 {
		return Cong + pont, x - a
	} else if second == "TIP" {
		return tip + Selecionado.Dica2 + pont, x - t
	} else {
		return respErr + pont, x - e
	}
}

func string3(third string, Selecionado Character, ponto2 int8) (resposta3 string, ponto3 int8) {
	x = ponto2
	if third == Selecionado.Nome || third == Selecionado.Nome2 || third == Selecionado.Nome3 {
		return Cong + pont, x - a
	} else if third == "TIP" {
		return tip + Selecionado.Dica3 + pont, x - t
	} else {
		return respErr + pont, x - e
	}
}

func string4(fourth string, Selecionado Character, ponto3 int8) (resposta4 string, ponto4 int8) {
	x = ponto3
	if fourth == Selecionado.Nome || fourth == Selecionado.Nome2 || fourth == Selecionado.Nome3 {
		return Cong + pont, x - a
	} else if fourth == "TIP" {
		return tip + Selecionado.Dica4 + pont, x - t
	} else {
		return respErr + pont, x - e
	}
}

func string5(fifth string, Selecionado Character, ponto4 int8) (resposta5 string, ponto5 int8) {
	x = ponto4
	if fifth == Selecionado.Nome || fifth == Selecionado.Nome2 || fifth == Selecionado.Nome3 {
		return Cong + pont, x - a
	} else if fifth == "TIP" {
		return " Não! Chega de dica! é melhor você beber um café! " + pont, 0
	} else {
		return respErr + pont, 0
	}
}

func main() {

	c1 := Character{}
	c1.SetChar(SeverChar())
	var revealChar string = "A resposta correta é: " + c1.Nome

	fmt.Println("O jogo vai começar.")
	fmt.Println("Digite um personagem para dar um palpite ou digite 'Tip' para pedir uma dica!")
	fmt.Println("Você começa com 6 pontos. Cada dica custa 1 ponto; Cada erro custa 2 pontos.")

	response := bufio.NewReader(os.Stdin)
	line, _ := response.ReadString('\r')
	line = strings.TrimSuffix(line, "\r")
	first := strings.ToUpper(line)
	fmt.Println(string1(first, c1))
	resposta1, ponto1 := string1(first, c1)
	if resposta1 == Cong+pont {
		os.Exit(0)
	}

	response2 := bufio.NewReader(os.Stdin)
	line2, _ := response2.ReadString('\r')
	line2 = strings.TrimSuffix(line2, "\r")
	second := strings.ToUpper(line2)
	fmt.Println(string2(second, c1, ponto1))
	resposta2, ponto2 := string2(second, c1, ponto1)
	if resposta2 == Cong+pont {
		os.Exit(0)
	}

	response3 := bufio.NewReader(os.Stdin)
	line3, _ := response3.ReadString('\r')
	line3 = strings.TrimSuffix(line3, "\r")
	third := strings.ToUpper(line3)
	fmt.Println(string3(third, c1, ponto2))
	resposta3, ponto3 := string3(third, c1, ponto2)
	if ponto3 <= 0 {
		fmt.Println(revealChar)
		os.Exit(0)
	}
	if resposta3 == Cong+pont {
		os.Exit(0)
	}

	response4 := bufio.NewReader(os.Stdin)
	line4, _ := response4.ReadString('\r')
	line4 = strings.TrimSuffix(line4, "\r")
	fourth := strings.ToUpper(line4)
	fmt.Println(string4(fourth, c1, ponto3))
	resposta4, ponto4 := string4(fourth, c1, ponto3)
	if ponto4 == 0 {
		fmt.Println(revealChar)
		os.Exit(0)
	} else if ponto4 == -1 {
		fmt.Println("Como você conseguiu ter pontuação negativa?!" + revealChar)
		os.Exit(0)
	}
	if resposta4 == Cong+pont {
		os.Exit(0)
	}

	response5 := bufio.NewReader(os.Stdin)
	line5, _ := response5.ReadString('\r')
	line5 = strings.TrimSuffix(line5, "\r")
	fifth := strings.ToUpper(line5)
	fmt.Println(string5(fifth, c1, ponto4))
	resposta5, ponto5 := string5(fifth, c1, ponto4)
	if ponto5 <= 0 {
		fmt.Println(revealChar)
		os.Exit(0)
	}
	if resposta5 == Cong+pont {
		os.Exit(0)
	}

}
