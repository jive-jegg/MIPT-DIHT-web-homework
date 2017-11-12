function deleteForm(e) {
	var elem = e.target.parentNode;
	elem.parentNode.removeChild(elem);
}

function addForm(txt) {
	var lis = document.createElement('li');
	document.getElementById('unordlist').appendChild(lis);
	var sp = document.createElement('span');
	lis.appendChild(sp);
	$(sp).text(txt);
	document.getElementById('add_task_input').value = "";
	var but = document.createElement('button');
	lis.appendChild(but);
	$(but).text("Удалить");
	but.addEventListener("click", deleteForm);
}

var inp = document.createElement('input');
inp.setAttribute("id", "add_task_input");
document.getElementById('root').appendChild(inp);

document.getElementById('root').appendChild(document.createElement('br'));
document.getElementById('root').appendChild(document.createElement('br'));			

var but_add = document.createElement('button');
but_add.setAttribute("id", "add_task");
$(but_add).text("Добавить элемент");
document.getElementById('root').appendChild(but_add);
but_add.addEventListener("click", function() { addForm($('#add_task_input').val()) });


var unlist = document.createElement('ul');
unlist.setAttribute('id', 'unordlist');
document.getElementById('root').appendChild(unlist);
addForm("Сделать задание #3 по web-программированию");
