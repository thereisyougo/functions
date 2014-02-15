import java.util.ArrayList;
import java.util.List;

public class CommonFun {
	public static final String trim(Object obj) {
		if (obj == null)
			return "";
		return obj.toString().trim();
	}

	/**
	 * 主要用于汉字超出指定长度后，分别截取为多个子片段
	 *
	 * @param limit
	 *            限制长度
	 * @param content
	 *            全文
	 * @return 片段
	 */
	public static final List<String> cut(int limit, String content) {
		List<String> list = new ArrayList<String>();
		cut0(limit, content, list);
		return list;
	}

	private static void cut0(int limit, String content, List<String> list) {
		byte[] bytes = content.getBytes();
		if (bytes.length <= limit) {
			list.add(content);
			return;
		}
		String section = new String(bytes, 0, limit);
		char lastChar = section.charAt(section.length() - 1);
		if (!Character.isLetter(lastChar) && !isChinese(lastChar))
			section = new String(bytes, 0, limit - 1);

		list.add(section);
		cut0(limit, content.substring(section.length()), list);
	}

	static boolean isAsciiPrintable(final char ch) {
		return ch >= 32 && ch < 127;
	}

	/**
	 * 根据Unicode编码判断中文汉字和符号
	 * 转自http://www.micmiu.com/
	 * @param c
	 * @return
	 */
	public static final boolean isChinese(char c) {
		Character.UnicodeBlock ub = Character.UnicodeBlock.of(c);
		if (ub == Character.UnicodeBlock.CJK_UNIFIED_IDEOGRAPHS
				|| ub == Character.UnicodeBlock.CJK_COMPATIBILITY_IDEOGRAPHS
				|| ub == Character.UnicodeBlock.CJK_UNIFIED_IDEOGRAPHS_EXTENSION_A
				|| ub == Character.UnicodeBlock.CJK_UNIFIED_IDEOGRAPHS_EXTENSION_B
				|| ub == Character.UnicodeBlock.CJK_SYMBOLS_AND_PUNCTUATION
				|| ub == Character.UnicodeBlock.HALFWIDTH_AND_FULLWIDTH_FORMS
				|| ub == Character.UnicodeBlock.GENERAL_PUNCTUATION) {
			return true;
		}
		return false;
	}

}
