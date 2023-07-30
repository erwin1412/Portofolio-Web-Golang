// // let nama = "erwin"

// // let age = 24

// // alert(`hello ${nama} umur ${age}`)

// //operation
// // let angka1 = 50
// // let angka2 = 100

// // alert (`hasil : ${angka1 + angka2} `)

// // let nilai = 70
// // if (nilai >= 90) {
// //     keterangan = "Grade A";
// //   } else if (nilai < 90 && nilai >= 80) {
// //     keterangan = "Grade B";
// //   } else {
// //     keterangan = "Grade C";
// //   }

// //   alert(keterangan)
function submitData(event) {

    

    
    let name =  document.getElementById("input-name").value
    let email =  document.getElementById("input-email").value
    let phone =  document.getElementById("input-phone").value
    let subject =  document.getElementById("input-subject").value
    let message =  document.getElementById("input-message").value

    let objectData = {
        name: name,
        email,
        phone,
        subject,
        message
    }

    let arrayData = [name, email, phone, subject, message]

    console.log(objectData)
    console.log(arrayData)


    console.log(name)
    console.log(email)
    console.log(phone)
    console.log(subject)
    console.log(message)
    if (name == '') {
        alert('name tidak boleh kosong')
        name = document.getElementById("input-name");
        name.focus();
        name.select();
        return ''
    }else if (email == '' ) {
         alert('email tidak boleh kosong')
         email = document.getElementById("input-email");
        email.focus();
        email.select();
        return ''
    }else if ( phone == '' ) {
         alert('phone tidak boleh kosong')
         phone = document.getElementById("input-phone");
        phone.focus();
        phone.select();
        return ''
    }else if (subject == '' ) {
         alert('subject tidak boleh kosong')
         subject = document.getElementById("input-subject");
        subject.focus();
        subject.select();
        return ''
    }else if (message == '' ) {
         alert('Message tidak boleh kosong')
         message = document.getElementById("input-message");
        message.focus();
        message.select();
        return ''
    }else{
    alert(`Nama Saya : ${name}\nEmail Saya : ${email}\nPhone Saya : ${phone}\nSubject Saya : ${subject}\nMessage Saya : ${message}\n` )


    const receiverEmail = "erwin@gmail.com"
    
    let a = document.createElement('a') 
    a.href = `mailto:${receiverEmail}?subject=${subject}&body=Halo nama saya ${name},\n${message}, silahkan kontak saya di nomor berikut : ${phone}`
    a.click();

    
    name =  document.getElementById("input-name").value = ''
    email =  document.getElementById("input-email").value = ''
    phone =  document.getElementById("input-phone").value = ''
    subject =  document.getElementById("input-subject").value = ''
    message =  document.getElementById("input-message").value = ''
    
}

}


// alternative
// const submitButton = document.getElementById('submit');

// submitButton.addEventListener('click', function(event) {
//   event.preventDefault(); // Prevent form submission

//   const fieldIds = ['input-name', 'input-email', 'input-phone', 'input-subject', 'input-message'];

//   for (const fieldId of fieldIds) {
//     const inputField = document.getElementById(fieldId);

//     if (inputField.value === '') {
//       alertField(`${getFieldLabel(fieldId)} field cannot be empty`, inputField);
//       return;
//     }
//   }
// });

// function getFieldLabel(fieldId) {
//   return fieldId;
// }

// function alertField(message, inputElement) {
//   alert(message);
//   inputElement.focus();
// }
