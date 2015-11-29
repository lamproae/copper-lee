function show_alert() {
	alert("++++++++++++++\n");
}

function show_confirm() {
	var result = confirm("是否删除");
	if (result) {
		alert("Delete success!\n");
	} else {
		alert("Do not delete");
	}
}

function show_prompt() {
	var value = prompt("Please input your name", "copper");
	if (value == null) {
		alert("Canceld\n");
	} else if (value == "") {
		alert("Empty name, input again!\n");
		show_prompt();
	} else {
		alert("你好, "+value);
	}
}

