import { useNavigate } from "react-router-dom";


interface Props {
}

const Home: React.FC<Props> = () => {

    const navigate = useNavigate();

    const goHome = () => {
    
        console.log("Goin home");

        navigate("/")
    };

    return(
        <>
            Home Page
            <button onClick={goHome}> Go Home </button>
        </>
    )
}
export default Home;