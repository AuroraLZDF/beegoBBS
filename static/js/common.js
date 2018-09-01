var userAgent = navigator.userAgent.toLowerCase();
var is_opera = userAgent.indexOf('opera') != -1 && opera.version();
var is_moz = (navigator.product == 'Gecko') && userAgent.substr(userAgent.indexOf('firefox') + 8, 3);
var is_ie = (userAgent.indexOf('msie') != -1 && !is_opera) && userAgent.substr(userAgent.indexOf('msie') + 5, 3);
var is_safari = (userAgent.indexOf('webkit') != -1 || userAgent.indexOf('safari') != -1);

function codeId(id) {
	_sTxt = $('#'+id).html();
	setCopy(_sTxt);
	$("#codetext").html('代码已复制到系统剪切板');
}

function setCopy(_sTxt){
	if(is_ie) {
		clipboardData.setData('Text',_sTxt);
		alert ("网址“"+_sTxt+"”\n已经复制到您的剪贴板中\n您可以使用Ctrl+V快捷键粘贴到需要的地方");
	} else {
		prompt("请复制网站地址:",_sTxt);
	}
}

function wh_size(_width, _height, zwidth, zheight) {
	var size = {};
	if(zwidth <= _width && zheight <= _height) {
		size.width= zwidth;
		size.height = zheight;

	} else {

		var r = zwidth/_width;
		var t = zheight/_height;
		if(parseInt(t*100) > parseInt(r*100)) {
			size.height = _height;
			size.width = zwidth/t;
		}else{
			size.width = _width;
				size.height = zheight/r;
		}
		/*if(zwidth >= _width && zheight >= _height){
			if(zwidth > zheight){
				size.width = _width;
				size.height = zheight/r;
			} else {
				size.height = _height;
				size.width = zwidth/t;
			}
		}else{
			if(zwidth >= _width &&  zheight <= _height) {
				size.width = _width;
				size.height = zheight/r;
			}else{
				size.height = _height;
				size.width = zwidth/t;
			}
		}*/
	}
	return size;



}

var showLayer = function ( selector,jsonStyle) {
    $(selector).show();
    if (jsonStyle && jsonStyle.position) {
        $(selector).css('position', jsonStyle.position);
    }
    if (jsonStyle && jsonStyle.width) {
        $(selector).css('width', jsonStyle.width);
    }
    if (jsonStyle && jsonStyle.height) {
        $(selector).css('height', jsonStyle.height);
    }
    if (jsonStyle && jsonStyle.top) {
        $(selector).css('margin-top', 0);
        $(selector).css('top', jsonStyle.top);
    }
    if (jsonStyle && jsonStyle.left) {
        $(selector).css('margin-left', 0);
        $(selector).css('left', jsonStyle.left);
    }
    $(".shade").show();
};

var showTimelocal = function(url){
	var i = 1;
	var f = {
		hid: function(){
			location.href = url;
		},
		dv: function(){
			i--;
		},
		ui: function(){
			if(i>0){
				f.dv();
				setTimeout(f.ui, 1000);
			} else {
				f.hid();
			}
		}
	};
	f.ui();
};


var showTimeHide = function(o, ob, m){
	var i = 2;
	var f = {
		hid: function(){
			o.hide();
			ob.html('');
		},
		dv: function(){
			if (m){
				ob.html(m);
			}
			o.show();
			i--;
		},
		ui: function(){
			if(i>0){
				f.dv();
				setTimeout(f.ui, 1000);
			} else {
				f.hid();
			}
		}
	};
	f.ui();
};
var fnTimeCountDown = function(o, m){
	var i = 2;
	var f = {
		hid: function(){
			o.hide('slow');
		},
		dv: function(){
			if (m){
				o.html(m);
			}
			o.show();
			i--;
		},
		ui: function(){
			if(i>0){
				f.dv();
				setTimeout(f.ui, 1000);
			} else {
				f.hid();
			}
		}
	};
	f.ui();
};


function subString(str, len, hasDot)
{
    var newLength = 0;
    var newStr = "";
    var chineseRegex = /[^\x00-\xff]/g;
    var singleChar = "";
    var strLength = str.replace(chineseRegex,"**").length;
    for(var i = 0;i < strLength;i++)
    {
        singleChar = str.charAt(i).toString();
        if(singleChar.match(chineseRegex) != null)
        {
            newLength += 2;
        }
        else
        {
            newLength++;
        }
        if(newLength > len)
        {
            break;
        }
        newStr += singleChar;
    }

    if(hasDot && strLength > len)
    {
        newStr += "...";
    }
    return newStr;
}
/****************************************
* 函数名称：IsDate
* 功能说明：构造函数
* 参    数：sDate:日期字符串
* 调用示列：
*           string sDate="2008-10-28";
*           IsDate(sDate);
*****************************************/
/// <summary>
/// 判断是否是日期
/// </summary>
/// <param name="sDate">日期字符串</param>
/// <returns>返回是否(bool)</returns>
function IsDate(sDate)
{
	var sRegex= /^(\d{4})-(\d{2})-(\d{2})$/;
	var bResult = sDate.match(sRegex);
	if(bResult==null)
	{
	    return   false;
	}
	else
	{
	    return   true;
	}
}

/****************************************
* 函数名称：IsNullEmpty
* 功能说明：判断字符串是否为空
* 参    数：str:空字符串
* 调用示列：
*           string str="";
*           IsNullEmpty(str);
*****************************************/
/// <summary>
///  判断字符串是否为空
/// </summary>
/// <param name="sNullOrEmpty">空字符串</param>
/// <returns>返回是否(bool)</returns>
function IsNullEmpty(sNullOrEmpty) {
    if (!sNullOrEmpty || sNullOrEmpty.length == '' || sNullOrEmpty.length <= 0) {
	    return false;
	}
	else
	{
	    return true;
	}
}

/****************************************
* 函数名称：IsCurrent
* 功能说明：判断是否是货币
* 参    数：sCurrent:货币字符串
* 调用示列：
*           string sCurrent="88888.00";
*           IsCurrent(sCurrent);
*****************************************/
/// <summary>
/// 判断是否是货币
/// </summary>
/// <param name="sCurrent">货币字符串</param>
/// <returns>返回是否(bool)</returns>
function IsCurrent(sCurrent)
{
	var bResult1=sCurrent.match("[^0-9.-]");
	var bResult2=sCurrent.match("[[0-9]*[.][0-9]*[.][0-9]*");
	var bResult3=sCurrent.match("[[0-9]*[-][0-9]*[-][0-9]");
	var bResult4=sCurrent.match("(^([-]|[.]|[-.]|[0-9])[0-9]*[.]*[0-9]+$)|(^([-]|[0-9])[0-9]*$)");
	if (bResult1!=null||bResult2!=null||bResult3!=null||bResult4==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}

/****************************************
* 函数名称：IsNumeric
* 功能说明：判断是否是数字
* 参    数：sNum:数字字符串
* 调用示列：
*           string sNum="88888";
*           IsNumeric(sNum);
*****************************************/
/// <summary>
/// 判断是否是数字
/// </summary>
/// <param name="sNum">数字字符串</param>
/// <returns>返回是否(bool)</returns>
function IsNumeric(sNum)
{
	var bResult=sNum.match("^(-|\\+)?\\d+(\\.\\d+)?$");
	if (bResult==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}

/****************************************
* 函数名称：IsUrl
* 功能说明：判断是否是URL
* 参    数：sUrl:URL字符串
* 调用示列：
*           string sUrl="http:\\www.sina.com.cn";
*           IsUrl(sUrl);
*****************************************/
/// <summary>
/// 判断是否是URL
/// </summary>
/// <param name="sUrl">URL字符串</param>
/// <returns>返回是否(bool)</returns>
function IsUrl(sUrl)
{
    sUrl = $.trim(sUrl);
    sUrl = sUrl.toLowerCase();
    if ((sUrl.substr(0, 7) != "http://")&&(sUrl.substr(0, 8) != "https://")) {
        sUrl = "http://"+sUrl;
    }

    var bResult = sUrl.match("http(s)?://([\\w-]+\\.)+[\\w-]+(/[\\w- ./?%&=]*)?");

    if (bResult == null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}

/****************************************
* 函数名称：IsMail
* 功能说明：判断是否是MAILL
* 参    数：sMail:MAIL字符串
* 调用示列：
*           string sMail="olivier@hdtworld.com";
*           IsMail(sMail);
*****************************************/
/// <summary>
/// 判断是否是MAIL
/// </summary>
/// <param name="sMail">MAIL字符串</param>
/// <returns>返回是否(bool)</returns>
function IsMail(sMail)
{
	var bResult=sMail.match("\\w+([-+.']\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*");
	if (bResult==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}

/****************************************
* 函数名称：IsPostCode
* 功能说明：判断是否是邮编
* 参    数：sPostCode:邮编字符串
* 调用示列：
*           string sPostCode="200001";
*           IsPostCode(sPostCode);
*****************************************/
/// <summary>
/// 判断是否是邮编
/// </summary>
/// <param name="sPostCode">邮编字符串</param>
/// <returns>返回是否(bool)</returns>
function IsPostCode(sPostCode)
{
	var bResult=sPostCode.match("^\\d{6}$");
	if (bResult==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}

/****************************************
* 函数名称：IsTelephone
* 功能说明：判断是否是电话号码
* 参    数：sTelephone:电话号码字符串
* 调用示列：
*           string sTelephone="66660000";
*           IsTelephone(sTelephone);
*****************************************/
/// <summary>
/// 判断是否是电话号码
/// </summary>
/// <param name="sTelephone">电话号码字符串</param>
/// <returns>返回是否(bool)</returns>
function IsTelephone(sTelephone)
{
	var bResult=sTelephone.match("^(\\(\\d{3}\\)|\\d{3}-)?\\d{8}$");
	if (bResult==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}

/****************************************
* 函数名称：IsMobile
* 功能说明：判断是否是手机号码
* 参    数：sMobile:手机号码字符串
* 调用示列：
*           string sMobile="1381101101101";
*           IsMobile(sMobile);
*****************************************/
/// <summary>
/// 判断是否是手机号码
/// </summary>
/// <param name="sMobile">手机号码字符串</param>
/// <returns>返回是否(bool)</returns>
function IsMobile(sMobile)
{
	var bResult=sMobile.match("^\\d{11}$");
	if (bResult==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}

/****************************************
* 函数名称：IsIDCard
* 功能说明：判断是否身份证
* 参    数：sIDCard:身份证字符串
* 调用示列：
*           string sIDCard="310106198210054xxx";
*           IsIDCard(sIDCard);
*****************************************/
/// <summary>
/// 判断是否是数字
/// </summary>
/// <param name="sSimNum">数字字符串</param>
/// <returns>返回是否(bool)</returns>
function IsIDCard(code)
{
	/*var bResult=sIDCard.match("^\\d{15}|\\d{18}$");
	if (bResult==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}  */
    var city={11:"北京",12:"天津",13:"河北",14:"山西",15:"内蒙古",21:"辽宁",22:"吉林",23:"黑龙江 ",31:"上海",32:"江苏",33:"浙江",34:"安徽",35:"福建",36:"江西",37:"山东",41:"河南",42:"湖北 ",43:"湖南",44:"广东",45:"广西",46:"海南",50:"重庆",51:"四川",52:"贵州",53:"云南",54:"西藏 ",61:"陕西",62:"甘肃",63:"青海",64:"宁夏",65:"新疆",71:"台湾",81:"香港",82:"澳门",91:"国外 "};
    var tip = "";
    var pass= true;

    if(!code || !/(18|19|20)?\d{2}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])/i.test(code)){
        tip = "身份证号格式错误";
        pass = false;
    }

    else if(!city[code.substr(0,2)]){
        tip = "地址编码错误";
        pass = false;
    }
    else{
        //18位身份证需要验证最后一位校验位
        if(code.length == 18){
            code = code.split('');
            //∑(ai×Wi)(mod 11)
            //加权因子
            var factor = [ 7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2 ];
            //校验位
            var parity = [ 1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2 ];
            var sum = 0;
            var ai = 0;
            var wi = 0;
            for (var i = 0; i < 17; i++)
            {
                ai = code[i];
                wi = factor[i];
                sum += ai * wi;
            }
            var last = parity[sum % 11];
            if(parity[sum % 11] != code[17]){
                tip = "校验位错误";
                pass =false;
            }
        }
    }
    if(!pass) alert(tip);
    return pass;
}

/****************************************
* 函数名称：IsCE
* 功能说明：判断是中英表达式
* 参    数：sCE:中英文表达式字符串
* 调用示列：
*           string sCE="HDT互动通";
*           IsCE(sCE);
*****************************************/
/// <summary>
/// 判断是中英表达式
/// </summary>
/// <param name="sCE">中英文表达式字符串</param>
/// <returns>返回是否(bool)</returns>
function IsCE(sCE)
{
	var bResult=sCE.match("^[a-zA-Z\\u4E00-\\u9FA5\\uF900-\\uFA2D]+$");
	if (bResult==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}

	/// <summary>
	/// 密码强度等级
	/// </summary>
	var pwdLevel;
	/// <summary>
	/// 密码中是否有字母
	/// </summary>
	var hasLetter;
	/// <summary>
	/// 密码中是否有大小写字母
	/// </summary>
	var hasULLetter;
	/// <summary>
	/// 密码中是否有数字
	/// </summary>
	var hasNumeric;
	/// <summary>
	/// 密码中是否有符号
	/// </summary>
	var hasSymbol;


/****************************************
* 函数名称：IsPasswordLevel
* 功能说明：判断密码强度
* 参    数：sPassword:密码字符串
* 调用示列：
*           string sPassword="abc123-_";
*           IsPasswordLevel(sPassword);
*****************************************/
/// <summary>
/// 判断密码强度
/// </summary>
/// <param name="sPassword">密码字符串</param>
/// <returns>返回强度等级(string)</returns>
function IsPasswordLevel(sPassword)
{
	pwdLevel = 0;
	if (sPassword == "" || sPassword == null)
	{
	    return "空密码";
	}
	else
	{
	    //判断密码长度
	    JugePwdLength(sPassword);
	    //判断字母
	    JugePwdLetter(sPassword);
	    //判断数字
	    JugePwdNumeric(sPassword);
	    //判断符号
	    JugeSymbol(sPassword);
	    //判断奖励
	    JugeAward();
	    //判断密码级别
	    //>= 90: 非常安全
	    //>= 80: 安全（Secure）
	    //>= 70: 非常强
	    //>= 60: 强（Strong）
	    //>= 50: 一般（Average）
	    //>= 25: 弱（Weak）
	    //>= 0: 非常弱
	    if (pwdLevel > 0)
	    {
	        if (pwdLevel > 25)
	        {
	            if (pwdLevel > 50)
	            {
	                if (pwdLevel > 60)
	                {
	                    if (pwdLevel > 70)
	                    {
	                        if (pwdLevel > 80)
	                        {
	                            if (pwdLevel > 90)
	                            {
	                                return "非常安全";
	                            }
	                            else
	                            {
	                                return "安全";
	                            }
	                        }
	                        else
	                        {
	                            return "非常强";
	                        }
	                    }
	                    else
	                    {
	                        return "强";
	                    }
	                }
	                else
	                {
	                    return "一般";
	                }
	            }
	            else
	            {
	                return "弱";
	            }
	        }
	        else
	        {
	            return "非常弱";
	        }
	    }
	    return "极其弱";
	}
}

/****************************************
* 函数名称：JugePwdlength
* 功能说明：判断密码字符串长度
* 参    数：str:字符串
* 调用示列：
*           string str="abc123-_";
*           JugePwdlength(str);
*****************************************/
/// <summary>
/// 判断密码字符串长度
/// </summary>
/// <param name="slength">密码字符串</param>
function JugePwdLength(sLength)
{
	var length = sLength.length;
	if (length <= 4)
	{
	    pwdLevel += 5;
	}
	else
	{
	    if (length <= 7)
	    {
	        pwdLevel += 10;
	    }
	    else
	    {
	        pwdLevel += 20;
	    }
	}
}

/****************************************
* 函数名称：JugePwdLetter
* 功能说明：判断密码强度是否有字符
* 参    数：str:字符串
* 调用示列：
*           string str="abc123-_";
*           JugePwdLetter(str);
*****************************************/
/// <summary>
/// 判断密码强度是否有字符
/// </summary>
/// <param name="sLetter">密码字符串</param>
function JugePwdLetter(sLetter)
{
	//0 分: 没有字母
	//10 分: 全都是小（大）写字母
	//20 分: 大小写混合字母
	//判断是否有字母
	var count = 0;
	var othercount = 0;
	var bLower=false, bUpper=false;
	for (var i = 0; i <= sLetter.length - 1; i++)
	{
	    //大小写字母的KEYCODE 65-90
	    if((sLetter.charCodeAt(i)>=65)&&(sLetter.charCodeAt(0)<=90))
	    {
	        count += 1;
	    }
	    //判断字符是否有大小写
	    if (sLetter.substr(i,1).match("[A-Z]"))
	    {
	        bUpper = true;
	    }
	    //判断字符是否有大小写
	    if (sLetter.substr(i,1).match("[a-z]"))
	    {
	        bLower = true;
	    }
	}
	if (count == 0)
	{
	    pwdLevel += 0;
	}
	else
	{
	    hasLetter = true;
	    if (bLower && bUpper)
	    {
	        pwdLevel += 20;
	    }
	    else
	    {
	        pwdLevel += 10;
	    }
	};
}

/****************************************
* 函数名称：JugePwdNumeric
* 功能说明：判断密码强度是否有数字
* 参    数：str:密码字符串
* 调用示列：
*           string str="abc123-_";
*           JugePwdNumeric(str);
*****************************************/
/// <summary>
/// 判断密码强度是否有数字
/// </summary>
/// <param name="str">密码字符串</param>
function JugePwdNumeric(sNum)
{
	//三、数字:
	//0 分: 没有数字
	//10 分: 1 个数字
	//20 分: 大于等于 3 个数字
	var count = 0;


	for (var i = 0; i <= sNum.length - 1; i++)
	{
	   //数字的KEYCODE 96-105
	   if((sNum.charCodeAt(i)>=96)&&(sNum.charCodeAt(0)<=105))
	    {
	        count += 1;
	    }
	}
	if (count == 0)
	{
	    pwdLevel += 0;
	}
	else
	{
	    hasNumeric = true;
	    if (count < 3)
	    {
	        pwdLevel += 10;
	    }
	    else
	    {
	        pwdLevel += 20;
	    }
	}
}

/****************************************
* 函数名称：JugeAward
* 功能说明：判断密码强度奖励
* 参    数：
* 调用示列：
*           JugeAward();
*****************************************/
/// <summary>
/// 判断密码强度奖励
function JugeAward()
{


	//五、奖励:
	//2 分: 字母和数字
	//3 分: 字母、数字和符号
	//5 分: 大小写字母、数字和符号
	if (hasLetter && hasNumeric)
	{
	    if (hasSymbol)
	    {
	        if (hasULLetter)
	        {
	            pwdLevel += 5;
	        }
	        else
	        {
	            pwdLevel += 3;
	        }
	    }
	    else
	    {
	        pwdLevel += 2;
	    }
	}
}


/****************************************
* 函数名称：JugeAward
* 功能说明：判断特定的符号
* 参    数：str:密码字符串
* 调用示列：
*           string str="abc123-_";
*           IsSymbol(str);
*****************************************/
/// <summary>
/// 判断特定的符号
/// </summary>
/// <param name="str">密码字符串</param>
/// <returns>返回是否(bool)</returns>
function IsSymbol(str)
{
	var bResult=str.match("[_]|[-]|[#]");
	if (bResult==null)
	{
	    return false;
	}
	else
	{
	    return true;
	}
}


/****************************************
* 函数名称：JugeSymbol
* 功能说明：判断是密码强度否有符号
* 参    数：str:密码字符串
* 调用示列：
*           string str="abc123-_";
*           JugeSymbol(str);
*****************************************/
/// <summary>
/// 判断是密码强度否有符号
/// </summary>
/// <param name="str">密码字符串</param>
function JugeSymbol(sSymbol)
{
	//四、符号:
	//0 分: 没有符号
	//10 分: 1 个符号
	//25 分: 大于 1 个符号


	var count = 0;
	var tmpstr = "";
	for (var i = 0; i <= sSymbol.length - 1; i++)
	{
	    tmpstr = sSymbol.substr(i, 1);
	    if (IsSymbol(tmpstr))
	    {
	        count += 1;
	    }
	}
	if (count == 0)
	{
	    pwdLevel += 0;
	}
	else
	{
	    hasSymbol = true;
	    if (count > 1)
	    {
	        pwdLevel += 25;
	    }
	    else
	    {
	        pwdLevel += 10;
	    }
	}
}

//Luhm校验规则：16位银行卡号（19位通用）:

// 1.将未带校验位的 15（或18）位卡号从右依次编号 1 到 15（18），位于奇数位号上的数字乘以 2。
// 2.将奇位乘积的个十位全部相加，再加上所有偶数位上的数字。
// 3.将加法和加上校验位能被 10 整除。

//bankno为银行卡号 banknoInfo为显示提示信息的DIV或其他控件
function luhmCheck(bankno) {
    if (!/^[0-9]*$/.test(bankno)) {
        return false;
    }
    return true;
    var lastNum=bankno.substr(bankno.length-1,1);//取出最后一位（与luhm进行比较）

    var first15Num=bankno.substr(0,bankno.length-1);//前15或18位
    var newArr=new Array();
    for(var i=first15Num.length-1;i>-1;i--){    //前15或18位倒序存进数组
        newArr.push(first15Num.substr(i,1));
    }

    var arrJiShu=new Array();  //奇数位*2的积 <9
    var arrJiShu2=new Array(); //奇数位*2的积 >9

    var arrOuShu=new Array();  //偶数位数组
    for(var j=0;j<newArr.length;j++){
        if((j+1)%2==1){//奇数位
            if(parseInt(newArr[j])*2<9)
                arrJiShu.push(parseInt(newArr[j])*2);
            else
                arrJiShu2.push(parseInt(newArr[j])*2);
        }
        else //偶数位
            arrOuShu.push(newArr[j]);
    }

    var jishu_child1=new Array();//奇数位*2 >9 的分割之后的数组个位数
    var jishu_child2=new Array();//奇数位*2 >9 的分割之后的数组十位数
    for(var h=0;h<arrJiShu2.length;h++){
        jishu_child1.push(parseInt(arrJiShu2[h])%10);
        jishu_child2.push(parseInt(arrJiShu2[h])/10);
    }

    var sumJiShu=0; //奇数位*2 < 9 的数组之和
    var sumOuShu=0; //偶数位数组之和
    var sumJiShuChild1=0; //奇数位*2 >9 的分割之后的数组个位数之和
    var sumJiShuChild2=0; //奇数位*2 >9 的分割之后的数组十位数之和
    var sumTotal=0;
    for(var m=0;m<arrJiShu.length;m++){
        sumJiShu=sumJiShu+parseInt(arrJiShu[m]);
    }

    for(var n=0;n<arrOuShu.length;n++){
        sumOuShu=sumOuShu+parseInt(arrOuShu[n]);
    }

    for(var p=0;p<jishu_child1.length;p++){
        sumJiShuChild1=sumJiShuChild1+parseInt(jishu_child1[p]);
        sumJiShuChild2=sumJiShuChild2+parseInt(jishu_child2[p]);
    }
    //计算总和
    sumTotal=parseInt(sumJiShu)+parseInt(sumOuShu)+parseInt(sumJiShuChild1)+parseInt(sumJiShuChild2);

    //计算Luhm值
    var k= parseInt(sumTotal)%10==0?10:parseInt(sumTotal)%10;
    var luhm= 10-k;
    if(lastNum==luhm){
        $("#banknoInfo").html("Luhm验证通过");
        return true;
    }
    else{
        $("#banknoInfo").html("银行卡号必须符合Luhm校验");
        return false;
    }
}

function luhmCheckNew(bankno) {
    if (!/^[0-9]*$/.test(bankno)) {
        return false;
    }
    return true;
    if (bankno.length < 16 || bankno.length > 19) {
        //$("#banknoInfo").html("银行卡号长度必须在16到19之间");
        return false;
    }
    var num = /^\d*$/;  //全数字
    if (!num.exec(bankno)) {
        //$("#banknoInfo").html("银行卡号必须全为数字");
        return false;
    }
    //开头6位
    var strBin = "10,18,30,35,37,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,58,60,62,65,68,69,84,87,88,94,95,98,99";
    if (strBin.indexOf(bankno.substring(0, 2)) == -1) {
        //$("#banknoInfo").html("银行卡号开头6位不符合规范");
        return false;
    }
    var lastNum = bankno.substr(bankno.length - 1, 1);//取出最后一位（与luhm进行比较）

    var first15Num = bankno.substr(0, bankno.length - 1);//前15或18位
    var newArr = new Array();
    for (var i = first15Num.length - 1; i > -1; i--) {    //前15或18位倒序存进数组
        newArr.push(first15Num.substr(i, 1));
    }
    var arrJiShu = new Array();  //奇数位*2的积 <9
    var arrJiShu2 = new Array(); //奇数位*2的积 >9

    var arrOuShu = new Array();  //偶数位数组
    for (var j = 0; j < newArr.length; j++) {
        if ((j + 1) % 2 == 1) {//奇数位
            if (parseInt(newArr[j]) * 2 < 9)
                arrJiShu.push(parseInt(newArr[j]) * 2);
            else
                arrJiShu2.push(parseInt(newArr[j]) * 2);
        }
        else //偶数位
            arrOuShu.push(newArr[j]);
    }

    var jishu_child1 = new Array();//奇数位*2 >9 的分割之后的数组个位数
    var jishu_child2 = new Array();//奇数位*2 >9 的分割之后的数组十位数
    for (var h = 0; h < arrJiShu2.length; h++) {
        jishu_child1.push(parseInt(arrJiShu2[h]) % 10);
        jishu_child2.push(parseInt(arrJiShu2[h]) / 10);
    }

    var sumJiShu = 0; //奇数位*2 < 9 的数组之和
    var sumOuShu = 0; //偶数位数组之和
    var sumJiShuChild1 = 0; //奇数位*2 >9 的分割之后的数组个位数之和
    var sumJiShuChild2 = 0; //奇数位*2 >9 的分割之后的数组十位数之和
    var sumTotal = 0;
    for (var m = 0; m < arrJiShu.length; m++) {
        sumJiShu = sumJiShu + parseInt(arrJiShu[m]);
    }

    for (var n = 0; n < arrOuShu.length; n++) {
        sumOuShu = sumOuShu + parseInt(arrOuShu[n]);
    }

    for (var p = 0; p < jishu_child1.length; p++) {
        sumJiShuChild1 = sumJiShuChild1 + parseInt(jishu_child1[p]);
        sumJiShuChild2 = sumJiShuChild2 + parseInt(jishu_child2[p]);
    }
    //计算总和
    sumTotal = parseInt(sumJiShu) + parseInt(sumOuShu) + parseInt(sumJiShuChild1) + parseInt(sumJiShuChild2);

    //计算Luhm值
    var k = parseInt(sumTotal) % 10 == 0 ? 10 : parseInt(sumTotal) % 10;
    var luhm = 10 - k;

    if (lastNum == luhm) {
        $("#banknoInfo").html("Luhm验证通过");
        return true;
    }
    else {
        $("#banknoInfo").html("银行卡号必须符合Luhm校验");
        return false;
    }
}



function getservers(gid){
    if (gid) {
        $.post("/game/server/getservers/", { gid: gid }, function (data) {
            var JSON = eval("(" + data + ")");
            var html = "";
            if (JSON) {
                for (var i = 0; i < JSON.length; i++) {
                    html += '<input type="checkbox" name="server_id[]" value="' + JSON[i].id + '" id="ck' + i + '">&nbsp;&nbsp;<label for="ck' + i + '">' + JSON[i].name + '</label>&nbsp;&nbsp;&nbsp;&nbsp;';
                }
            }
            $(".server_list").html(html);
        });
    }
}

//刷新验证码
function refresh_yzcode() {
    $('.yz').attr("src", "/login/index/code?_t=" + Math.random());
}

//兼容placeholder
function hasPlaceholderSupport() {
    return 'placeholder' in document.createElement('input');
}

/*************************
 *函数名称：select_game
 *功能说明：通过选择游戏名称刷新服务器列表，并使刚刚选择的服务器处于选中状态
 *参    数：gid:游戏gid;server_id:服务器id
 *调用示例：
 var gid = ''
 var server_id = '';
 *			select_game(gid,server_id)
 *************************/
function select_game(gid, _server_id) {
	$("input[name=gid]").val(gid);
	var options = "<option value=''>请选择服务器</option>";
	var server_element = $("select[name=server_id]");
	$.getJSON("/index/script/getServer?_t="+Math.random(), { gid: gid }, function (data) {
		var server_list = data;

		for (var i = 0; i < server_list.length; i++) {
			var ut = '';
			if(_server_id == server_list[i].id) {
				ut='selected';
			}
			options += "<option value='" + server_list[i].id + "' gid='"+gid+"' "+ut+"   >" + server_list[i].name + "</option>";
		}
		server_element.html(options);
	});
}

var PlaceHolder = {
    _support: (function () {
        return 'placeholder' in document.createElement('input');
    })(),

    //提示文字的样式，需要在页面中其他位置定义
    className: 'abc',

    init: function () {
        if (!PlaceHolder._support) {
            //未对textarea处理，需要的自己加上
            var inputs = document.getElementsByTagName('input');
            PlaceHolder.create(inputs);
        }
    },

    create: function (inputs) {
        var input;
        if (!inputs.length) {
            inputs = [inputs];
        }
        for (var i = 0, length = inputs.length; i < length; i++) {
            input = inputs[i];
            if (!PlaceHolder._support && input.attributes && input.attributes.placeholder) {
                PlaceHolder._setValue(input);
                input.addEventListener('focus', function (e) {
                    if (this.value === this.attributes.placeholder.nodeValue) {
                        this.value = '';
                        this.className = '';
                    }
                }, false);
                input.addEventListener('blur', function (e) {
                    if (this.value === '') {
                        PlaceHolder._setValue(this);
                    }
                }, false);
            }
        }
    },

    _setValue: function (input) {
        input.value = input.attributes.placeholder.nodeValue;
        input.className = PlaceHolder.className;
    }
};

//页面初始化时对所有input做初始化
PlaceHolder.init();
//或者单独设置某个元素
$(function () {
    $("#fromSearch").submit(function () {
		if($.banFormRepeatSubmit.isBan(this)){
			return false;
		}
        $(this).find("input[placeholder]").each(function () {
            if ($(this).val() == $(this).attr("placeholder")) {
                $(this).val("");
            }
        });
		$.banFormRepeatSubmit.removeBan(this);
    }).addClass($.banFormRepeatSubmit.otherClass());
    /*$("input[placeholder]").each(function () {
        PlaceHolder.create($(this)[0]);
    });*/
});



(function(jQ){
	var banClass = 'mzqBanClass', otherClass = 'mzqFreedomClass';
	jQ.banFormRepeatSubmit = function (){
		$('form').each(function(){
			var _thisjQ = $(this);
			var events = jQuery._data(this).events;
			if(!events || !events.submit){
				bindSubmit(_thisjQ);
			}else if(events.submit){
				testOtherForm(_thisjQ);
			}else{
				testMsg(_thisjQ, 'is waring');
			}
		});


		function testMsg(formjQuery, msg){
			console && console.log && console.log('form name: ' + formjQuery.attr('name') + ' id: #' + formjQuery.attr('id') + ' ' + msg);
		}

		function testOtherForm(formjQuery){
			if(!formjQuery.hasClass(otherClass)){
				testMsg(formjQuery, 'not has ' + otherClass);
			}
		}

		function bindSubmit(formjQuery){
			if(formjQuery.hasClass(otherClass)){
				return ;
			}
			formjQuery.addClass(otherClass);
			formjQuery.submit(function (){
				if(formjQuery.hasClass(banClass)){
					return false;
				}
				formjQuery.addClass(banClass);
				setTimeout(function(){
					formjQuery.removeClass(banClass);
				}, 3000);
			});
		}
	}

	jQ.banFormRepeatSubmit.otherClass = function(){
		return otherClass;
	}

	jQ.banFormRepeatSubmit.banClass = function(){
		return banClass;
	}

	jQ.banFormRepeatSubmit.isBan = function(obj){
		var banForm = obj.jquery ? obj : $(obj);
		//console.log(banForm)
		if(banForm.hasClass(banClass)){
			return true;
		}
		banForm.addClass(banClass);
		return false;
	}

	jQ.banFormRepeatSubmit.removeBan = function(obj){
		var banForm = obj.jquery ? obj : $(obj);
		if(banForm.hasClass(banClass)){
			banForm.removeClass(banClass);
		}
	}

})(jQuery);