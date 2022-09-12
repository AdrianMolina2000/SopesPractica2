import {useEffect, useState} from 'react';
import axios from "axios";
import './App.css';
import React from 'react';
import {Pie} from 'react-chartjs-2';
import {Chart, ArcElement} from 'chart.js'
import 'chart.js/auto';
import 'bootstrap/dist/css/bootstrap.min.css';

Chart.register(ArcElement);

function App(){
  const [ram, setRam] = useState({});
  const [cpu, setCpu] = useState({});
  const [det, setDetalle] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:5000/getRam').then(response => {
      setRam(JSON.parse(response.data[0].valor));
    }).catch(err => console.log(err))
    
    axios.get('http://localhost:5000/getCpu').then(response => {
      setCpu(JSON.parse(response.data[0].valor));
    }).catch(err => console.log(err))
  }, []);

  console.log(ram);
  console.log(cpu);

  var ramPer = (ram.total - ram.free) * 100 / ram.total;
  const data = {
    labels: ['Usado %', 'Total'],
    datasets: [{
      data: [ramPer, 100],
      backgroundColor: ['#3e95cd', '#8e5ea2'],
    }]
  };

  const options = {
    title: {
      display: true,
      text: 'Porcentaje de uso de la RAM'
    },
    responsive: true,
  }

  const columnas = [
    {
      name: 'name',
      selector: 'name',
      sortable: true,
    },
    {
      name: 'pid',
      selector: 'pid',
      sortable: true,
    },
    {
      name: 'ram',
      selector: 'ram',
      sortable: true,
    },
    {
      name: 'state',
      selector: 'state',
      sortable: true,
    },
    {
      name: 'user',
      selector: 'user',
      sortable: true,
    },
  ];

  function getTabla() {
    let filas = [];
    let index = 1;
    cpu.data.map((item) => {
      if(item.state == 1 || 1026){
        item.state = "Sleeping"
      }else if(item.state == 0){
        item.state = "Running"
      }else if(item.state == 4){
        item.state = "Zombie"
      }else if(item.state == 8){
        item.state = "Stopped"
      }
      
      filas.push(<tr>
        <td>{index}</td>
        <td>{item.pid}</td>
        <td>{item.name}</td>
        <td>{item.user}</td>
        <td>{item.state}</td>
        <td>{item.ram}</td>
      </tr>)
      index++;
    })
    
    setDetalle(filas);
  }
  
  
  return (
    <div className="App">
    <h1>Uso Actual de la Ram</h1>
      <div className='box'>
        <Pie data={data} options={options} />
      </div>
      
      <div>
        <table class="default">
        <tr>
          <td>No.</td>
          <td>Pid</td>
          <td>Name</td>
          <td>User</td>
          <td>State</td>
          <td>Ram</td>
        </tr>

        {
          det
        }
      </table>
      </div>
      
      <button onClick={() => {
        getTabla();
      }}
      >
        Actualizar
      </button>
    </div>
  );
}
export default App;