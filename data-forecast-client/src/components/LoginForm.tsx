import React, {ChangeEvent, FormEvent, useState, useEffect} from "react";
 
interface Props {
}

const LoginForm: React.FC<Props> = () => {
    return (
        <form onSubmit={handleSubmit}>
            <input type="text" placeholder="name"/>
            <input type="text" placeholder="address" />
            <button> Submit </button>
        </form>
    )
};

const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    console.log("Trying to submit the form");
};

const onNameChange = (e: ChangeEvent<HTMLInputElement>) => {
    //setName(e.target.value);
};

export default LoginForm;