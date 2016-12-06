public class Keyprac {
    public static void main(String args[]) {
        OthersKeyy1 k1 = new OthersKeyy1();

        String key = "lots0of0fun0in0l";
        byte[] bytes = key.getBytes();
        String hex = javax.xml.bind.DatatypeConverter.printHexBinary(bytes);
        System.out.println(hex);
        char[][] matrix = new char[4][32];
        char[][] matrix2 = new char[4][1];

        int i, j, n = 0, k, p = 0;
        for (j = 0; j < 4; j++) {
            for (k = n; k < n + 8; k++) {
                matrix[j][k] = hex.charAt(k);
                System.out.print(matrix[j][k] + " ");
            }
            n = k;
            System.out.println();
        }

        for (k = 1; k <= 4; k++) {
            matrix2[p][0] = matrix[p][7 * k + p];
            System.out.println(matrix2[p][0]);
            p++;
        }
        k1.rot(matrix2);
    }
}


    
