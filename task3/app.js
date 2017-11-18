function addList(list, text) {
    var par = $('<li>');
    par.append('<span>' + text + '</span>');
    par.append(
            $('<button class="deleter">')
        .text("Удалить")
        .click(function () {
        $(this).parent().remove();
    }));
    $("#task_list").append(par);
}
$(document).ready(function () {
    var list = $("#root").append('<ul id="task_list"></ul>');
    addList(list, "Сделать задание #3 по web-программированию");
    $('#root').append('<input type="input" id="add_task_input">');
    $('#root').append('<input type="button" id="add_task" value="OK">');
    $("#add_task").click(function () {
        addList(list, $('#add_task_input').val());
    });

});

