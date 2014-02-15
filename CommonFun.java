import java.util.ArrayList;
import java.util.List;


public class CommonFun {
	
	public static final String trim(Object obj) {
		if (obj == null) return "";
		return obj.toString().trim();
	}
	
	/**
	 * 主要用于汉字超出指定长度后，分别截取为多个子片段
	 * @param limit 限制长度
	 * @param content 全文
	 * @return 片段
	 */
	public static final List<String> cut(int limit, String content) {
		List<String> list = new ArrayList<String>();
		cut0(limit, content, list);
		return list;
	}

	private static void cut0(int limit, String content, List<String> list) {
		int last = 0;
		byte[] bytes = content.getBytes();
		if (bytes.length <= limit) {
			list.add(content);
			return;
		}
		if (isAsciiPrintable((char) bytes[limit - 1]))
			last = limit;
		else
			last = limit - 1;
		String section = new String(bytes, 0, last);
		list.add(section);
		cut0(limit, content.substring(section.length()), list);
	}

	static boolean isAsciiPrintable(final char ch) {
        	return ch >= 32 && ch < 127;
    	}

}
