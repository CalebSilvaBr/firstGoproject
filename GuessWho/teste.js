var dc = "Sou da DC"

const charA = {
    names: ["Magneto", "Maguineto", "Magineto"],
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

console.log(thissortedChar)
console.log(thissortedChar.names)
/*
let sorted = getRandomNumber(0, group.length-1)

let sortedChar = group[sorted]

console.log(sortedChar)
*/