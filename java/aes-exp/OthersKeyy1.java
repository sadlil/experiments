public class OthersKeyy1 {
	char[][] matrix3= new char[4][1];

	public void rot(char[][] matrix2) {
		matrix3[3][0] = matrix2[0][0];
		for(int i=0; i<3; i++) {
			matrix3[i][0] = matrix2[i+1][0];
		}

		for (int i=0; i<matrix2.length; i++) {
			System.out.print(matrix3[i][0]);
			System.out.println();
		}
	}
}
