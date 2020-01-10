import React, {useState} from 'react';
import logo from './utils/logo.jpg';
import './App.css';
import "jquery";
import "jquery-ui";
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
    const [viewChanged, setViewChanged] = useState(0);

    const handleViewChange = (event, view) => {
        setViewChanged(prevState => prevState === 0 ? 1 : 0);
        setView(view);
    };
    const handleUserIdChange = (event) => setUserId(event.target.value);

    const handleLogoClick = () => {
      setView('Bank');
      setUserId("");
    };
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" onClick={handleLogoClick}/>
                <p className="welcome-top"> Question Bank </p>
                <div className="user-input">
                    <label className="input-label"> Enter the user Id: </label>
                    <input className="userId-text" type="text" value={userId} onChange={handleUserIdChange}/>
                </div>
            </header>
            <div className="App-body">
                <AppBar position="static" className="tab-view">
                    <Tabs value={view} onChange={handleViewChange} centered="true">
                        <Tab label="Bank App" title="Open bank app" value="Bank"/>
                        <Tab label="User App" title="Open user app" value="Customer"/>
                        <Tab label="User Validation" title="Open user validation app" value="Validation"/>
                    </Tabs>
                </AppBar>
                {userId && view === "Bank" && <BankView userId={userId} viewChanged={viewChanged}/>}
                {userId && view === "Customer" && <CustomerView userId={userId} viewChanged={viewChanged}/>}
                {userId && view === "Validation" && <ValidationView userId={userId} viewChanged={viewChanged}/>}
            </div>
        </div>
    );
}

export default App;
