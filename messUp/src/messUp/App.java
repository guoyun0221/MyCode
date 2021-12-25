package messUp;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.util.Scanner;

/**
 * I wanted to use this program for any type of file,
 * I test it on images first,
 * but the images processed by this procedure can't displayed anymore,
 * maybe it's because after reorder bytes, the rgb is over 255.
 * now it's just work on text file.
 * @author GuoYun
 */
public class App {
	
	// since the program only used on text file, 
	// the header is not needed.
	// static final int HEADER_LENGTH = 64;
	// static final int REAR_LENGTH = 16;
	
	Scanner scan = new Scanner(System.in);
	
	public static void main(String[] args) {
		App app = new App();
		boolean success = app.messUp();
		System.out.println("success: " + success);
	}
	
	/**
	 * reorder bytes of the file
	 * @return true: success; false: failed
	 */
	public boolean messUp() {
		// get the file into byte array
		byte[] buffer = getFile();
		if (buffer == null) {
			return false;
		}
		// reorder the bytes 
		// there may be more than one way to do this
		swapAdjacentTwoBytes(buffer);
		// output the file 
		boolean success = output(buffer);
		if (!success) {
			return false;
		}
		scan.close();
		return true;
	}
	
	public byte[] getFile() {
		System.out.println("input file name: ");
		String fileName = scan.nextLine();
		File f = new File(fileName);
		if (!(f.exists())) {
			System.out.println("File doesn't exist\n");
			return null;
		}
		long len = f.length();
		if (len > Integer.MAX_VALUE) {
			System.out.println("The file is too large\n");
			return null;
		}
		byte[] bytes = new byte[(int)f.length()];
		try (FileInputStream in = new FileInputStream(fileName)) {
			in.read(bytes);
		} catch (Exception e) {
			e.printStackTrace();
		}
		return bytes;
	}
	
	public void swapAdjacentTwoBytes(byte[] bytes) {
		for (int i = 0, len = bytes.length - 1; i < len; i += 2) {
			byte t = bytes[i];
			bytes[i] = bytes[i + 1];
			bytes[i + 1] = t;
		}
	}
	
	/**
	 * 
	 * @param bytes to output file
	 * @return success(true) or not(false);
	 */
	public boolean output(byte[] bytes) {
		System.out.println("output file name: ");
		String fileName = scan.nextLine();
		
		File f = new File(fileName);

		try (FileOutputStream out = new FileOutputStream(fileName)) {
			f.createNewFile();
			out.write(bytes);
		} catch (Exception e) {
			e.printStackTrace();
		}
		
		return true;
	}
}
