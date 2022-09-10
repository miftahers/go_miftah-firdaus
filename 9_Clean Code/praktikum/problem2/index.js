// PROBLEM 2 - REWRITE
/*
class kendaraan {
	var totalroda = 0;
	var kecepatanperjam = 0;
}

class mobil extends kendaraan {
	void berjalan () {
		tambahkecepatan(10);
	}
}

tambahkecepatan(var kecepatanbaru) {
	kecepatanperjam = kecepatanperjam + kecepatanbaru;
}

void main () {
	mobilcepat = new mobil();
	mobilcepat.berjalan();
	mobilcepat.berjalan();
	mobilcepat.berjalan();

	mobillamban = new mobil();
	mobillamban.berjalan();
}
*/

class Kendaraan {
	roda = 0;
	kecepatan = 0;
}

class Mobil extends Kendaraan {
	berjalan(){
		this.kecepatan += 10;
	}
}

function main() {
	mobilCepat = new Mobil()
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLambat = new Mobil()
	mobilLambat.berjalan()

	console.log(mobilCepat.kecepatan)
	console.log(mobilLambat.kecepatan)
}