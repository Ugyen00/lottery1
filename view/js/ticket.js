function resetform() {
    document.getElementById("tid").value = "";
    document.getElementById("fname").value = "";
    document.getElementById("lname").value = "";
    document.getElementById("phone").value = "";
}

function GetNo() {
let generatedNumbers = [];

function generateRandomNumber() {
  return Math.floor(Math.random() * (999 - 100 + 1)) + 100;
}

function isNumberGenerated(number) {
  return generatedNumbers.includes(number);
}

while (generatedNumbers.length < 900) {
  let randomNumber = generateRandomNumber();
  if (!isNumberGenerated(randomNumber)) {
    generatedNumbers.push(randomNumber);
  }
}
var genNo = generatedNumbers[Math.floor(Math.random() * generatedNumbers.length)];
document.getElementById("tid").value = genNo;
}

function buyTicket() {
    var data = {
        tikid : parseInt(document.getElementById("tid").value),
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        phone : parseInt(document.getElementById("phone").value),
    }
    fetch('/ticket', {
        method : "POST",
        body: JSON.stringify(data),
        headers: {"Content-type":"application/json; charset=UTF-8"}
    }).then(response => {
        if (response.status == 201) {
            fetch('/ticket/'+ data.tikid)
                .then(response => response.text())
                .then(data => showTicket(data))
        } else {
            throw new Error(response.statusText)
        }
    }).catch(e => {
        alert(e)
    })
    resetform();
    // location.reload();
}

function newRow(ticket) {
    var table = document.getElementById("TicketList");
    var row = table.insertRow(table.length);
    var td=[]
    for(i=0; i<table.rows[0].cells.length; i++){
        td[i] = row.insertCell(i);
        }
    td[0].innerHTML = ticket.tikid;
    td[1].innerHTML = ticket.fname;
    td[2].innerHTML = ticket.lname;
    td[3].innerHTML = ticket.phone;
    if (document.cookie != ""){
      td[4].innerHTML = '<input type="button" onclick="deleteTicket(this)"value="Delete" id="button-1">';
      td[5].innerHTML = '<input type="button" onclick="updateTicket(this)"value="Edit" id="button-2">';
    }
    
}

function showTicket(data) {
    const ticket = JSON.parse(data)
    newRow(ticket)
}

function showTickets(data) {
    const tickets = JSON.parse(data)
    tickets.forEach(tik => {
        newRow(tik)
    })
}

window.onload = function(data) {
    fetch('/tickets')
        .then(response => response.text())
        .then(data => showTickets(data));
}

var selectedRow= null;
function deleteTicket(r) {
    if (confirm('Are you sure you want to DELETE this?')) {
        selectedRow = r.parentElement.parentElement;
        tid = selectedRow.cells[0].innerHTML;

        fetch('/ticket/' + tid, {
            method: "DELETE",
            headers: { "Content-type": "application/json; charset=UTF-8" }
        });
        var rowIndex = selectedRow.rowIndex;
        if (rowIndex > 0) {
            document.getElementById("TicketList").deleteRow(rowIndex);
        }
        selectedRow = null;
    }
}

var selectedRow = null;
function updateTicket(r) {
  selectedRow = r.parentElement.parentElement;
  document.getElementById("tid").value = selectedRow.cells[0].innerHTML;
  document.getElementById('fname').value = selectedRow.cells[1].innerHTML;
  document.getElementById('lname').value = selectedRow.cells[2].innerHTML;
  document.getElementById('phone').value = selectedRow.cells[3].innerHTML;

  var btn = document.getElementById("button-add");
  tik = selectedRow.cells[0].innerHTML;
  if(btn) {
    btn.innerHTML = "Update";
    btn.setAttribute("onclick", "update(tik)");
  }
}
function getFormData(){
  var data = {
    tikid : parseInt(document.getElementById("tid").value),
    fname : document.getElementById("fname").value,
    lname : document.getElementById("lname").value,
    phone : parseInt(document.getElementById("phone").value)
  }
  return data
}
function update(tid) {
  var newData = getFormData()
  console.log(tid,"hiiii")
  fetch('/ticket/'+tid, {
    method: "PUT",
    body: JSON.stringify(newData),
    headers: {"Content-type": "application/json; charset=UTF-8"}
  }).then (res => {
    if (res.ok) {
      selectedRow.cells[0].innerHTML = newData.tkid;
      selectedRow.cells[1].innerHTML = newData.fname;
      selectedRow.cells[2].innerHTML = newData.lname;
      selectedRow.cells[3].innerHTML = newData.phone;

      var button = document.getElementById("button-add");
      button.innerHTML = "Add";
      button.setAttribute("onclick", "addStudent()");
      selectedRow = null;
      resetform();
    }else {
      alert("Server: Update request error.")
      throw new Error(res.statusText)
    }
  })
  .catch(e =>{
    alert(e)
  })
}


//time
// let countDownDate = new Date("May 15, 2023 14:25:00").getTime();
// let container = document.querySelectorAll(".tcontainer");

// let countdownInterval = setInterval(function () {
//     let now = new Date().getTime();
//     let distance = countDownDate - now;
//     let days = Math.floor(distance / (1000 * 60 * 60 * 24));
//     let hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
//     let minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
//     let seconds = Math.floor((distance % (1000 * 60)) / 1000);

//     container[0].childNodes[1].innerHTML = days;
//     container[0].childNodes[3].innerHTML = 'days';
//     container[1].childNodes[1].innerHTML = hours;
//     container[1].childNodes[3].innerHTML = 'hours';
//     container[2].childNodes[1].innerHTML = minutes;
//     container[2].childNodes[3].innerHTML = 'minutes';
//     container[3].childNodes[1].innerHTML = seconds;
//     container[3].childNodes[3].innerHTML = 'seconds';

//     if (distance <= 0) {
//         clearInterval(countdownInterval); // Stop the interval
//         let win =  Math.floor(Math.random() * 900 + 100);
//         alert("Countdown has ended! Winner is Lottery Number: " + win); // Display an alert
//     }
// }, 1000);
let container = document.querySelectorAll(".tcontainer");
let countdownInterval;
let loopCount = 0;

function startCountdown() {
  let countDownDate = new Date().getTime() + 86400000; // Set initial countdown to 24 hours from current time
  loopCount++;

  countdownInterval = setInterval(function() {
    let now = new Date().getTime();
    let distance = countDownDate - now;
    let days = Math.floor(distance / (1000 * 60 * 60 * 24));
    let hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    let minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
    let seconds = Math.floor((distance % (1000 * 60)) / 1000);

    container[0].childNodes[1].innerHTML = days;
    container[0].childNodes[3].innerHTML = 'days';
    container[1].childNodes[1].innerHTML = hours;
    container[1].childNodes[3].innerHTML = 'hours';
    container[2].childNodes[1].innerHTML = minutes;
    container[2].childNodes[3].innerHTML = 'minutes';
    container[3].childNodes[1].innerHTML = seconds;
    container[3].childNodes[3].innerHTML = 'seconds';

    if (distance <= 0) {
      clearInterval(countdownInterval);
      if (loopCount >= 2) {
        let win = Math.floor(Math.random() * 900 + 100);
        alert("Countdown has ended! Winner is Lottery Number: " + win);
        loopCount = 0; // Reset loop count
      } else {
        startCountdown(); // Start the countdown again
      }
    }
  }, 1000);
}
// Start the initial countdown
startCountdown();
