import Web3 from 'web3';
import Crud from '../build/contracts/Crud.json';

let web3;
let crud;

const initWeb3 = () => {
  return new Promise((resolve, reject) => {
    if(typeof window.ethereum !== 'undefined') {
      const web3 = new Web3(window.ethereum);
      window.ethereum.enable()
        .then(() => {
          resolve(
            new Web3(window.ethereum)
          );
        })
        .catch(e => {
          reject(e);
        });
      return;
    }
    if(typeof window.web3 !== 'undefined') {
      return resolve(
        new Web3(window.web3.currentProvider)
      );
    }
    resolve(new Web3('http://localhost:9545'));
  });
};

const initContract = () => {
  const deploymentKey = Object.keys(Crud.networks)[0];
  return new web3.eth.Contract(
    Crud.abi, 
    Crud
      .networks[deploymentKey]
      .address
  );
};

const initApp = () => {
  const $create = document.getElementById('create');
  const $createResult = document.getElementById('create-result');
  const $read = document.getElementById('read');
  const $readResult = document.getElementById('read-result');
  let accounts = [];
  var $a=12;
  if($a==12){
  	document.getElementById('create').style.display ="none";
  }else{
    document.getElementById('read').style.display ="none";
  }



  web3.eth.getAccounts()
  .then(_accounts => {
    accounts = _accounts;
  });

  $create.addEventListener('submit', (e) => {
    e.preventDefault();
    const name = e.target.elements[0].value;
    crud.methods.create("m1","p1",name,"imagenlol").send({from: accounts[0]})
    .then(result => {
      $createResult.innerHTML = `Registro añadido satisfactoriamente!`;
    })
    .catch(_e => {
      $createResult.innerHTML = `Error! No se añadió el registro`;
    });
  });

  $read.addEventListener('submit', (e) => {
    e.preventDefault();
    const id = e.target.elements[0].value;
    crud.methods.read(id).call()
    .then(result => {
      //for (var i=0;i<result.length;i++) {
     //   $readResult.innerHTML += `${result[i]}`;
     // }
      //$readResult.innerHTML = `${result[0]}`;
      var html = "<table border='5' width='100%'><col style='width:10%'><col style='width:10%'> <col style='width:40%'><col style='width:40%'>";
      html+="<thead> <tr>";
    html+="<td style='text-align:center'>"+"Código de Médico"+"</td>";
        html+="<td style='text-align:center'>"+"CI de Paciente"+"</td>";
            html+="<td style='text-align:center'>"+"Registro"+"</td>";
                html+="<td style='text-align:center'>"+"Imagen"+"</td> </tr></thead><tbody>";
for (var i = 0; i < result.length; i++) {
    html+="<tr>";
    html+="<td style='text-align:center'>"+`${result[i].codmedico}`+"</td>";
        html+="<td style='text-align:center'>"+`${result[i].codpaciente}`+"</td>";
            html+="<td style='text-align:center'>"+`${result[i].registro}`+"</td>";
                html+="<td style='text-align:center'>"+`${result[i].imagen}`+"</td>";
    html+="</tr>";

}
html+="</tbody></table>";
//document.getElementById("box").innerHTML = html;
$readResult.innerHTML = `${html}`;
    })
    .catch(_e => {
      $readResult.innerHTML = `Error! No se pudo leer el historial`;
    });
  });

};

document.addEventListener('DOMContentLoaded', () => {
  initWeb3()
    .then(_web3 => {
      web3 = _web3;
      crud = initContract();
      initApp(); 
    })
    .catch(e => console.log(e.message));
});
