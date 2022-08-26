class TennisGame2 {
  constructor(player1Name, player2Name) {
    this.player1 = this.createPlayer(player1Name);
    this.player2 = this.createPlayer(player2Name);
  }

  createPlayer(name) {
    return {points: 0, name: name, res: ""}
  }

  getScoreValueForPoint(number) {
    var scoreType = {
      0: "Love",
      1: "Fifteen",
      2: "Thirty", 
      3: "Forty"
    }
    return scoreType[number]
  }

  getScore() {
    var score = "";

    if (this.player1.points === this.player2.points) {
      score = this.getScoreValueForPoint(this.player1.points)
      score += "-All";  
    } else {
      this.player1.res = this.getScoreValueForPoint(this.player1.points)
      this.player2.res = this.getScoreValueForPoint(this.player2.points)
      score = this.player1.res + "-" + this.player2.res;
    }

    var scoreDiff = Math.abs(this.player1.points - this.player2.points)

    if (scoreDiff == 0 && this.player1.points > 2)
      score = "Deuce";

    if (scoreDiff == 1 && (this.player1.points > 3 || this.player2.points > 3)) {
      score = this.player1.points > this.player2.points ? "Advantage player1" : "Advantage player2"
    }

    if (scoreDiff >= 2 && (this.player1.points >= 4 || this.player2.points >= 4)) {
      score = this.player1.points > this.player2.points ? "Win for player1" : "Win for player2"
    }

    return score;
  };

  wonPoint(player) {
    player === "player1" ? this.player1.points++ : this.player2.points++
  };
}
if (typeof window === "undefined") {
  module.exports = TennisGame2;
}