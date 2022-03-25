import React, {ChangeEvent, FormEvent, useState, useEffect} from "react";
import { useNavigate } from "react-router-dom";
import LoginForm from '../components/LoginForm';
import logo from '../logo.svg';
import '../App.css';
// import { useHistory } from 'react-router-dom'
// import useRouter from 'hooks/useRouter';

interface Props {
}

// const {push} = useRouter();

const Login: React.FC<Props> = () => {

    // const history = useHistory();
    const navigate = useNavigate();
    

    const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
    
        console.log("Trying to submit the form");
    
        navigate("/Home")

        // push("/login");  
    };

    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p>
                DataForecast App
                </p>
                <LoginForm />
            </header>
        </div>
    )
};



const onNameChange = (e: ChangeEvent<HTMLInputElement>) => {
    //setName(e.target.value);
};

export default Login;