import React, {ChangeEvent, FormEvent, useState, useEffect} from "react";
 
interface Props {
}

const LoginForm: React.FC<Props> = () => {
    return (
        <form>
            <input type="text" placeholder="name"/>
            <input type="text" placeholder="address" />
            <button> Submit </button>
        </form>
    )
};

export default LoginForm;