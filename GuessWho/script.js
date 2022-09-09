var dc = 'Sou da DC';
var marvel = 'Sou da Marvel';
var mitologia = 'Sou da mitologia';
var x = [0];
var pontos = [6];

const charA = {
    names: ["MAGNETO", "MAGUINETO", "MAGINETO"],
    chartips: ["sou forte", "controlo metais", "minha filha é a Wanda"]
    }
    
const charB = {
        names: ["SUPERMAN", "SUPER MAN", "SUPER HOMEM"],
        chartips: [dc, " Não sou da terra. ", " Meu principal uniforme é azul. ", " Meu nome também é Kalel. "],
    }
    
const charC = {
        names: ["SHAZAM", "CHAZAM", "XAZAM"],
        chartips: [dc, " Sou uma criança ", " Me transformo em adulto. ", " Meu uniforme leva tem um raio amarelo no peito."],
    }

function getRandomNumber(min, max){

    let step1 = max - min;
    let step2 = Math.random()*step1;
    let result = Math.floor(step2) + min;
    
        return result
    }
    
let group = [charA, charB, charC]
let sortedChar = getRandomNumber(0, group.length)
let thissortedChar = group[sortedChar]

function compute(){

    var principal = document.getElementById("principal").value;
    var principalC = principal.toUpperCase()

    if (thissortedChar.names.includes(principalC)) {
        pontos -= 0
        document.getElementById("result").innerHTML = "Seu palpite é: " + principal + ". Você acertou! Seus pontos:" + pontos
    } else if (principal == "tip") {
        pontos -= 1
        document.getElementById("result").innerHTML = "Sua dica: " + thissortedChar.chartips[x] + pontos
        x++
    } else {
        pontos -= 2
        document.getElementById("result").innerHTML = "Errou! Seus pontos: " + pontos
    }
    
    if (pontos <= 0) {        
        document.getElementById("result").innerHTML = "Você perdeu! Seus pontos: " + pontos
    }
    return [pontos, x]
}

/*

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

*/