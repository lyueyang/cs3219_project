import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import CreateForm from './CreateForm';
import ReadForm from './ReadForm';
import UpdateForm from './UpdateForm';
import DeleteForm from './DeleteForm';
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
  <React.StrictMode>
    <h1 className="ml-3 mt-3">CRUD UI</h1>
    <div className = "container">
      <div className="row">
        <div className="col">
          <CreateForm />
        </div>
        <div className="col">
          <UpdateForm />
        </div>
      </div>
      <div className="row">
        <div className="col">
          <ReadForm />
        </div>
        <div className="col">
          <DeleteForm />
        </div>
      </div>
    </div>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
