
const courseSelectBox = document.querySelector("#courseDropdown");

const heaings = document.querySelector(".card-title");

let div = document.createElement('div');
div.className = "alert";
div.innerHTML = "<strong>aHi there!</strong> You've read an important message.";

document.querySelector(".card-body").append(div);

heaings.innerHTML = "Iqbal"


console.log(courseSelectBox);


var xhttp = new XMLHttpRequest();
xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
        // Typical action to be performed when the document is ready:
        let response = JSON.parse(xhttp.response);
        let optionHTML = `<option selected value"0">Select Course</option>`

        console.log(response,response.length);

        for(var i = 0; i < response.length; i++) {

            // var course = response[i];
            var opt = document.createElement('option');
            opt.innerHTML = opt.value = response[i].name
            courseSelectBox.append(opt);

            // optionHTML = `
            //     <option value="${course.name}">${course.name}</option>
            // `
            // courseSelectBox.append(optionHTML);
        }
    }
};
xhttp.open("POST", "/get-cources", true);
xhttp.send();