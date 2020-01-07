import React, {useState} from 'react';
import logo from './utils/logo.jpg';
import './App.css';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import BankView from "./view-renderer/bankView";
import CustomerView from "./view-renderer/customerView";
import ValidationView from "./view-renderer/validationView";
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
    const [view, setView] = useState('Bank');
    const [userId, setUserId] = useState("");
    const [questions, setQuestions] = useState(null);
    const [answers, setAnswers] = useState();

    function handleViewChange(event, view) {
        setView(view);
    }

    function handleUserIdChange(event) {
        setUserId(event.target.value)
    }

    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo"/>
                    <p className="welcome-top">Welcome to the Question Manager</p>
                <div className="user-input">
                    <label className="input-label">Enter the user Id: </label>
                    <input className="userId-text" type="text" value={userId} onChange={handleUserIdChange}/>
                </div>
            </header>
            <div className="App-body">
                {/*<div>*/}
                {/*    <p>Welcome to the Question Manager</p>*/}
                {/*    <label>Enter the user Id: </label>*/}
                {/*    <input className="userId-text" type="text" value={userId} onChange={handleUserIdChange}/>*/}
                {/*</div>*/}
                <AppBar position="static" className="tab-view">
                    <Tabs value={view} onChange={handleViewChange} centered="true">
                        <Tab label="Bank App" title="Open bank app" value="Bank"/>
                        <Tab label="Customer App" title="Open Customer app" value="Customer"/>
                        <Tab label="Customer Validation" title="Open customer validation app" value="Validation"/>
                    </Tabs>
                </AppBar>

                {userId && view === "Bank" && <BankView userId={userId}/>}
                {userId && view === "Customer" && <CustomerView userId={userId}/>}
                {userId && view === "Validation" && <ValidationView userId={userId}/>}
            </div>
        </div>
    );
}

export default App;
