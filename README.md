functions
=========

常用方法

# Markdown 支持以下这些符号前面加上反斜杠来帮助插入普通的符号：

> \   反斜线   
> `   反引号  
> \*   星号  
> _   底线  
> {}  花括号  
> []  方括号  
> ()  括弧  
> \#   井字号  
> \+   加号  
> \-   减号  
> .   英文句点  
> !   惊叹号  

# auto link
<http://example.com/>

# picture
![Alt text](/path/to/img.jpg)

# code
Use the `printf()` function.
# 插入一个引号 尾部插入两个以上空白可以插入一个
``There is a literal backtick (`) here.``    
A single backtick in a code span: `` ` ``    
A backtick-delimited string in a code span: `` `foo` ``    
# 使用<>不用转义
Please don't use any `<blink>` tags.     

# emphasis
*single asterisks*

_single underscores_

**double asterisks**

__double underscores__

un*frigging*believable

\*this text is surrounded by literal asterisks\*

# reference
This is [an example][id317] reference-style link.


[id317]: http://example.com/  "Optional Title Here"
[id318]: http://example.com/longish/path/to/resource/here
    "Optional Title Here"

# link
See my [About](/about/) page for details.

# 隐式链接标记
Visit [Daring Fireball][] for more information.


[Daring Fireball]: http://daringfireball.net/

# multilink
I get 10 times more traffic from [Google][1] than from
[Yahoo][2] or [MSN][3].

  [1]: http://google.com/        "Google"
  [2]: http://search.yahoo.com/  "Yahoo Search"
  [3]: http://search.msn.com/    "MSN Search"


I get 10 times more traffic from [Google][] than from
[Yahoo][] or [MSN][].

  [google]: http://google.com/        "Google"
  [yahoo]: http://search.yahoo.com/  "Yahoo Search"
  [msn]: http://search.msn.com/    "MSN Search"

This is [an example](http://example.com/ "Title") inline link.

[This link](http://example.net/) has no title attribute.

1986. What a great season.

1986\. What a great season.

> This is a blockquote with two paragraphs. Lorem ipsum dolor sit amet,
consectetuer adipiscing elit. Aliquam hendrerit mi posuere lectus.
Vestibulum enim wisi, viverra nec, fringilla in, laoreet vitae, risus.

> Donec sit amet nisl. Aliquam semper ipsum sit amet velit. Suspendisse
id sem consectetuer libero luctus adipiscing.

> This is the first level of quoting.
>
> > This is nested blockquote.
>
> Back to the first level.

> ## This is a header.
> 
> 1.   This is the first list item.
> 2.   This is the second list item.
> 
> Here's some example code:
>
>       return shell_exec("echo $input | $markdown_script");

*   Red
*   Green
*   Blue

+   Red
+   Green
+   Blue

-   Red
-   Green
-   Blue

1.  Bird
2.  McHale
3.  Parish

<ol>
<li>Bird</li>
<li>McHale</li>
<li>Parish</li>
</ol>



*   Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
    Aliquam hendrerit mi posuere lectus. Vestibulum enim wisi,
    viverra nec, fringilla in, laoreet vitae, risus.
*   Donec sit amet nisl. Aliquam semper ipsum sit amet velit.
    Suspendisse id sem consectetuer libero luctus adipiscing.

*   Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
Aliquam hendrerit mi posuere lectus. Vestibulum enim wisi,
viverra nec, fringilla in, laoreet vitae, risus.
*   Donec sit amet nisl. Aliquam semper ipsum sit amet velit.
Suspendisse id sem consectetuer libero luctus adipiscing.

# p
*   Bird
*   Magic

# list
1.  This is a list item with two paragraphs. Lorem ipsum dolor
    sit amet, consectetuer adipiscing elit. Aliquam hendrerit
    mi posuere lectus.

    Vestibulum enim wisi, viverra nec, fringilla in, laoreet
    vitae, risus. Donec sit amet nisl. Aliquam semper ipsum
    sit amet velit.

2.  Suspendisse id sem consectetuer libero luctus adipiscing.

# slack
*   This is a list item with two paragraphs.

    This is the second paragraph in the list item. You're
only required to indent the first line. Lorem ipsum dolor
sit amet, consectetuer adipiscing elit.

*   Another item in the same list.

# constract
*   A list item with a blockquote:

    > This is a blockquote
    > inside a list item.

# code area
*   A list item with a code block:

        <code goes here>
