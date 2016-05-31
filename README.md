##基本符号
*,-,+ 3个符号效果都一样，这3个符号被称为 Markdown符号
空白行表示另起一个段落
`是表示inline代码，tab是用来标记 代码段，分别对应html的code，pre标签

##换行
   
单一段落( <p>) 用一个空白行
连续两个空格 会变成一个 <br>
连续3个符号，然后是空行，表示 hr横线
##标题
生成h1--h6,在文字前面加上 1--6个# 来实现
文字加粗是通过 文字左右各两个符号
##引用
在第一行加上 “>”和一个空格，表示代码引用，还可以嵌套
##列表
这个是markdown文件的主要表示方式，主题要点化

使用*,+,-加上一个空格来表示
可以支持嵌套
有序列表用 数字+英文点+空格来表示
列表内容很长，不需要手工输入换行符，css控制段落的宽度，会自动的缩放的
##链接
直接写 [锚文本](url "可选的title")
引用 先定义 [ref_name]:url，然后在需要写入url的地方， 这样使用[锚文本][ref_name]，通常的ref_name一般用数字表示，这样显得专业
简写url：用尖括号包裹url 
这样生成的url锚文本就是url本身
##插入图片
一行表示: ![alt_text](url "可选的title")
引用表示法: ![alt_text][id],预先定义 [id]:url "可选title"
直接使用<img>标签，这样可以指定图片的大小尺寸
##特殊符号
用\来转义，表示文本中的markdown符号
可以在文本种直接使用html标签，但是要注意在使用的时候，前后加上空行
文本前后各加一个符号，表示斜体
## 插入数学公式
>自从使用Markdown以来，就开始一直使用Markdown+Github在写文章，整理自己的所学所思。本文亦是通过这种方式完成的。

然而，Markdown自由书写的特性很好，唯独遇到数学公式时就要煞费苦心——每次都是先使用Latex书写(在线的Latex编辑器参考[1])，然后保存为图片，使用img标签进行引用，当公式很多的时候稍显复杂。

本文的方法使用html的语法，调用[1]的公式生成API，在线生成Latex数学公式，免去将公式保存为图片的麻烦。当然，弊端也是有的，公式太多，可能会造成刷新比一般的网页慢一些。

方法一：使用Google Chart的服务器

<img src="http://chart.googleapis.com/chart?cht=tx&chl= 在此插入Latex公式" style="border:none;">
一个例子，

<img src="http://chart.googleapis.com/chart?cht=tx&chl=\Large x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}" style="border:none;">
公式显示结果为：



适用了下，Google Chart服务器的响应速度还可以，但据说可能复杂一些的Latex公式可能无法解析（参考[2]）。

方法二：使用forkosh服务器

forkosh上提供了关于Latex公式的一份简短而很有用的帮助，参考[1]和[3].

使用forkosh插入公式的方法是

<img src="http://www.forkosh.com/mathtex.cgi? 在此处插入Latex公式">
给个例子，

<img src="http://www.forkosh.com/mathtex.cgi? \Large x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}">
显示结果为：



因为网页插入公式的原理是调用“某某网站的服务器”动态生成的，所有保证公式正常显示的前提是该网址能一直存在着为我等小生做些小小的服务。forkosh我是用了快2年了，一直很好，推荐！

方法三：使用MathJax引擎

大家都看过Stackoverflow上的公式吧，漂亮，其生成的不是图片。这就要用到MathJax引擎，在Markdown中添加MathJax引擎也很简单，

<script type="text/javascript" src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=default"></script>
然后，再使用Tex写公式。$$公式$$表示行间公式，本来Tex中使用\(公式\)表示行内公式，但因为Markdown中\是转义字符，所以在Markdown中输入行内公式使用\\(公式\\)，如下代码：

$$x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}$$
\\(x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}\\)
分别显示结果（行间公式）：

x=−b±b2−4ac‾‾‾‾‾‾‾‾√2a
行内公式：

x=−b±b2−4ac√2a
不信，你可以试一下，在公式上还可以使用鼠标右键操作。
