class TennisGame2 {
  constructor(player1Name, player2Name) {
    this.player1 = this.createPlayer(player1Name);
    this.player2 = this.createPlayer(player2Name);
  }

  createPlayer(name) {
    return {points: 0, name: name}
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

  checkIfGameWon(score) {
    var scoreDiff = Math.abs(this.player1.points - this.player2.points)

    if (scoreDiff == 0 && this.player1.points > 2)
      return "Deuce";

    if (scoreDiff == 1 && (this.player1.points > 3 || this.player2.points > 3)) {
      return this.player1.points > this.player2.points ? "Advantage player1" : "Advantage player2"
    }

    if (scoreDiff >= 2 && (this.player1.points >= 4 || this.player2.points >= 4)) {
      return this.player1.points > this.player2.points ? "Win for player1" : "Win for player2"
    }
    return score
  }

  getScore() {
    var score = "";

    if (this.player1.points === this.player2.points) {
      score = this.getScoreValueForPoint(this.player1.points) + "-All";
    } else {
      score = this.getScoreValueForPoint(this.player1.points) + "-" + this.getScoreValueForPoint(this.player2.points);
    }

    return this.checkIfGameWon(score)
  };

  wonPoint(player) {
    player === this.player1.name ? this.player1.points++ : this.player2.points++
  };
}

if (typeof window === "undefined") {
  module.exports = TennisGame2;
}