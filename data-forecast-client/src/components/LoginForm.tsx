import React, {ChangeEvent, FormEvent, useState, useEffect} from "react";
import { useNavigate } from "react-router-dom";
// import { useHistory } from 'react-router-dom'

import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import ButtonToolbar from 'react-bootstrap/ButtonToolbar';
// import { LinkContainer } from 'react-router-bootstrap';


interface Props {
}

const LoginForm: React.FC<Props> = () => {

    const navigate = useNavigate();
    // const history = useHistory();

    const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
    
        console.log("Trying to submit the form! and navigate");

        navigate("/home")
    };

    const goHome = () => {
    
        console.log("Goin home");

        navigate("/")
    };

    return (
        <>
            <form onSubmit={handleSubmit}>
                <input type="text" placeholder="name"/>
                <input type="text" placeholder="address" />
                {/* <button onClick={goHome}> Go Home </button> */}
                <button> Submit </button>
            </form>
            <ButtonToolbar className="custom-btn-toolbar">
                <Button>Home</Button>
                <Button>About</Button>
                <Button>Users</Button>
            </ButtonToolbar>
        </>
    )
};



const onNameChange = (e: ChangeEvent<HTMLInputElement>) => {
    //setName(e.target.value);
};

export default LoginForm;