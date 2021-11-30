
// var previewCources = [{
//     "name" : "Iqbal"
// }];

// var xhr = new XMLHttpRequest();
// xhr.onload = function(e) {
//     if (this.status == 200) {
//         // Create a new Blob object using the response data of the onload object
//         var blob = new Blob([this.response], {type: 'text/plain'});
//         //Create a link element, hide it, direct it towards the blob, and then 'click' it programatically
//         let a = document.createElement("a");
//         a.style = "display: none";
//         document.body.appendChild(a);
//         //Create a DOMString representing the blob and point the link element towards it
//         let url = window.URL.createObjectURL(blob);
//         a.href = url;
//         a.download = 'myFile.txt';
//         //programatically click the link to trigger the download
//         a.click();
//         //release the reference to the file by revoking the Object URL
//         window.URL.revokeObjectURL(url);
//     }else{
//         //deal with your error state here
//     }
// };
// xhr.open("GET", "/save-file", true);
// xhr.send();

// const input = document.getElementById('fileInput2');

// const inputHandler = function(evt) {

//     console.log("HELERE");

//     var reader = new FileReader();
//     reader.readAsText(file.files[0], "UTF-8");
//     reader.onload = function (evt) {
//         document.getElementById("fileContents").innerHTML = evt.target.result;
//     }
//     reader.onerror = function (evt) {
//         document.getElementById("fileContents").innerHTML = "error reading file";
//     }
// }

// input.addEventListener('click', inputHandler);

// var file = document.getElementById("fileInput").files[0];
// if (file) {
    
// }