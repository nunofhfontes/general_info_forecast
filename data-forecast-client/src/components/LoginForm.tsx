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
                <div className="input-group mb-3">
                    <span className="input-group-text" id="basic-addon1">@</span>
                    <input type="text" className="form-control" placeholder="Username" 
                    aria-label="Username" aria-describedby="basic-addon1" />
                </div>
                <div className="input-group mb-3">
                    <span className="input-group-text" id="inputGroup-sizing-default">Password</span>
                    <input type="password" className="form-control" aria-label="Sizing example input" aria-describedby="inputGroup-sizing-default" />
                </div>
                <ButtonToolbar className="custom-btn-toolbar d-flex justify-content-center">
                    <Button type="submit">Submit</Button>
                </ButtonToolbar>
            </form>
        </>
    )
};



const onNameChange = (e: ChangeEvent<HTMLInputElement>) => {
    //setName(e.target.value);
};

export default LoginForm;