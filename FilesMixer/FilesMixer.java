import java.io.File;
import java.io.IOException;
import java.io.FileReader;
import java.io.FileReader;
import java.io.BufferedReader;
import java.io.FileWriter;
import java.io.BufferedWriter;
import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * does not support multi line keyword.
 */
public class FilesMixer{

    private static final String FILE_SEPARATOR_START = "*********file-separator-start*********";
    private static final String FILE_SEPARATOR_END = "*********file-separator-end*********";
    private static final String FILE_NAME = "filename:";
    
    private String fileSeparator = System.getProperty("file.separator");

    private String keyword;
    private String extension;
    
    private FileWriter fileWriter;
    private BufferedWriter bw;
    
    public static void main(String[] args) {
    	
    	FilesMixer fm = new FilesMixer();
    	
    	// mix
//    	fm.mix("main(", "c");
    	
    	// separate
    	fm.separate("MixedFile230110051126.txt");
    }

    /**
     * Recursively traverse the current directory to find target files
     * and mix them together to one file. 
     * @param keyword specify keyword in target file content, 
     * 	if this param is null, that means only filter file extension
     * @param extension specify target file extension, 
     * 	if this param is null, that means only filter keyword
     */
    public void mix(String keyword, String extension){
    	if (keyword == null && extension == null) {
    		System.out.println("You need to specify at least one param");
    	}
    	
    	this.keyword = keyword;
    	this.extension = extension;
    	
    	try {
    		Date date = new Date();
    		SimpleDateFormat dateFormat= new SimpleDateFormat("yyMMddhhmmss");
    		this.fileWriter = new FileWriter("MixedFile" + dateFormat.format(date) + ".txt");
            this.bw = new BufferedWriter(fileWriter);
    	} catch (IOException e) {
    		e.printStackTrace();
    	}

    	File root = new File(".");
    	doMix(root);
    	
    	try {
    		this.bw.close();
    		this.fileWriter.close();
    	} catch (Exception ex) {
			ex.printStackTrace();
		}
    }
    
    /**
     * read mixed file and generate directory and files according to it.
     * @param filename name of the mixed file 
     */
    public void separate(String filename){
    	try (BufferedReader br = new BufferedReader(new FileReader(filename))) {
    		// current line position. 0: header;  1: mixed file content
    		int status = -1;
    		String operationFile = null;
    		
			String line;
			while ((line = br.readLine()) != null) {
//				System.out.println("line:" + line);
				// chec line
				if (line.equals(FILE_SEPARATOR_START)) {
					// header start line
					// set status to 0
					status = 0;
				} else if (line.equals(FILE_SEPARATOR_END)) {
					// header end line
					// set status to 1
					status = 1;
				} else if (status == 0) {
					// in header
					if (line.indexOf(FILE_NAME) != -1) {
						// filename info 
						// mark filename
						operationFile = line.substring(line.indexOf(FILE_NAME) + 9);
						System.out.println("file: " + operationFile);
						// create file (and directory, maybe)
						touchFile(new File(operationFile));
					} else {
						System.out.println("parse header error");
					}
				} else if (status == 1) {
					// in file content
					// copy file
					this.fileWriter = new FileWriter(operationFile);
		            this.bw = new BufferedWriter(fileWriter);
		            
		            // write current line
		            bw.write(line + "\n");
		            
		            // write other lines
		            while ((line = br.readLine()) != null) {
		            	// file content end
		            	if (line.equals(FILE_SEPARATOR_START)) {
		            		status = 0;
		            		break;
		            	}
		            	bw.write(line + "\n");
		            }
		            
		            this.bw.close();
		            this.fileWriter.close();
				} else if (status == -1) {
					// init status
				} else {
					System.out.println("error: unexpected situation");
				}
			}
    		
    	} catch (Exception e) {
    		e.printStackTrace();
    	}
    }
    
    /**
     * @param file current operation file or directory
     */
    private void doMix(File file) {
    	File[] fs = file.listFiles();
    	for (File f: fs) {
    		if (f.isDirectory()) {
    			// is directory, search sub directory of it
    			doMix(f);
    		} else if (f.isFile()) {
    			// is file, do operation
//    			System.out.println("filename: " + f);
    			if (fileNeeded(f)) {
    				// copy this file to target mixed file
    				System.out.println("file: " + f);
    				copyFile(f);
    			}
    		}
    	}
    }
    
    private boolean fileNeeded(File f) {
    	
    	// check extension
    	if (this.extension != null) {
    		String filePath = f.toString();
        	String fileName = filePath.substring(filePath.lastIndexOf(this.fileSeparator) + 1);
        	String extension = fileName.substring(fileName.lastIndexOf(".") + 1);
//        	System.out.println(extension);
        	if (!this.extension.equals(extension)) {
        		return false;
        	}
    	}
    	
    	if (this.keyword == null) {
    		return true;
    	} else {
    		// check keyword
    		boolean pass = false;
    		
    		try (BufferedReader br = new BufferedReader(new FileReader(f))) {
                String line;
                while ((line = br.readLine()) != null) {
//                    System.out.println(line);
                    if (line.indexOf(this.keyword) != -1) {
//                    	System.out.println("keyword pass" + line.indexOf(this.keyword));
                    	pass = true;
                    }
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
    		
    		return pass;
    	}
    	
    }
    
    private void copyFile(File f) {
    	
    	try (BufferedReader br = new BufferedReader(new FileReader(f))) {
    		// write header
    		bw.write(FILE_SEPARATOR_START + "\n");
    		bw.write(FILE_NAME + f.toString() + "\n");
    		bw.write(FILE_SEPARATOR_END + "\n");
    		// copy file
    		String line;
            while ((line = br.readLine()) != null) {
                bw.write(line + "\n");
            }
    		
    	} catch (Exception e) {
    		e.printStackTrace();
    	} 
    	
    }
    
    private void touchFile(File file) {
    	if (file.exists()) {
            System.out.println("error: File exists");
        } else {
        	// if need to create parent directory
            if (!file.getParentFile().exists()) {
                file.getParentFile().mkdirs();
            }
            try {
                file.createNewFile();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
