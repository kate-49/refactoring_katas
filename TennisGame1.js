class TennisGame1 {
  constructor(player1Name, player2Name) {
    this.player1 = this.createPlayer(player1Name)
    this.player2 = this.createPlayer(player2Name)
  }

  createPlayer(name) {
    return {name: name, score: 0}
  }

  wonPoint(playerName) {
    playerName === this.player1.name? this.player1.score += 1: this.player2.score += 1;
  };

  checkForEqualScore(score) {
    var scoreType = {
      0: "Love-All",
      1: "Fifteen-All",
      2: "Thirty-All", 
    }
    return (scoreType[score] || "Deuce");
  }

  checkIfGameWon(player1Score, player2Score) {
    var minusResult = player1Score - player2Score;
    if (minusResult === -1) return "Advantage player2";
    else if (minusResult === 1) return "Advantage player1";
    else if (minusResult >= 2) return "Win for player1";
    else return "Win for player2";
  }

  calculateIndividualScores(score, tempScore) {
    var scoreType = {
      0: "Love",
      1: "Fifteen",
      2: "Thirty", 
      3: "Forty"
    }
    return score += scoreType[tempScore]
  }

  getScore() {
    var score = "";
    if (this.player1.score === this.player2.score) {
      score = this.checkForEqualScore(this.player1.score)
    } else if (this.player1.score >= 4 || this.player2.score >= 4) {
      score = this.checkIfGameWon(this.player1.score, this.player2.score)
    } else {
      score = this.calculateIndividualScores(score, this.player1.score)
      score += "-";
      score = this.calculateIndividualScores(score, this.player2.score)
    }
    return score;
  };
}
if (typeof window === "undefined") {
    module.exports = TennisGame1;
}