function addList(list, text) {
    var par = list.append('<li></li>').find('li').last();
    par.append('<span>' + text + '</span>');
    par.append('<button class="deleter">').text("Удалить");
    $(".deleter").click(function () {
        $(this).parent().remove();
    });
}
$(document).ready(function () {
    var list = $("#root").append('<ul></ul>').find('ul');
    addList(list, "Сделать задание #3 по web-программированию");
    $('#root').append('<input type="input" id="add_task_input">');
    $('#root').append('<input type="button" id="add_task" value="OK">');
    $("#add_task").click(function () {
        addList(list, $('#add_task_input').val());
    });

});

