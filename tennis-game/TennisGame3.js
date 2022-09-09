const points = ["Love", "Fifteen", "Thirty", "Forty"];

class TennisGame3 {
    constructor(player1Name, player2Name) {
        this.p1 = {name: player1Name, points: 0}
        this.p2 = {name: player2Name, points: 0}
    }

    calculatePoints() {
        const score = points[this.p1.points];
        return (this.p1.points == this.p2.points) ? score + "-All" : score + "-" + points[this.p2.points];
    }

    calculateWinner() {
        var winner = this.p1.points > this.p2.points ? this.p1.name : this.p2.name;
        return ((this.p1.points - this.p2.points) * (this.p1.points - this.p2.points) == 1) ? "Advantage " + winner : "Win for " + winner;
    }

    getScore() {
        if ((this.p1.points < 4 && this.p2.points < 4) && (this.p1.points + this.p2.points < 6)) {
            return this.calculatePoints();
        } else {
            if (this.p1.points == this.p2.points)
                return "Deuce";
            return this.calculateWinner()
        }
    }

    wonPoint(playerName) {
        playerName === this.p1.name ? this.p1.points += 1 : this.p2.points += 1;
    }
}

if (typeof window === "undefined") {
    module.exports = TennisGame3;
}
