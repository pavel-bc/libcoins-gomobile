package libcoins;

import com.sun.jna.Library;
import com.sun.jna.Native;

public class Main {
  static GoCqmTransformer GO_CQM_TRANSFORMER;
  static {
    String pwd = System.getProperty("user.dir");
    String lib = pwd + "/libcoins.so";
    GO_CQM_TRANSFORMER = (GoCqmTransformer) Native.loadLibrary(lib, GoCqmTransformer.class);
  }

  public interface GoCqmTransformer extends Library {
    String Do();
  }

  public static void main(String[] args) {
    String total = GO_CQM_TRANSFORMER.Do();
    System.out.println(total);
  }
}
