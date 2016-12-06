public class OthersKeyy1 {
	char[][] matrix3= new char[4][1];

	public void rot(char[][] matrix2) {

        for (int i=0; i<matrix2.length; i++) {
            for (int j=0; j<matrix2[i].length; j++) {
                System.out.print(" " + matrix2[i][j]);
            }
            System.out.println();
        }

		for(int i=0; i<3; i++) {
			matrix3[i][0] = matrix2[i+1][0];
			matrix3[3][0] = matrix2[0][0];
			System.out.println(matrix3[i][0]);
		}
	}
}
