import { Container, Nav, Navbar } from "react-bootstrap";
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
        <div className="homePage vh-100 grey-background">
            <Navbar className="mb-2" bg="dark" variant="dark">
                <Container>
                    <Navbar.Brand>
                        <img
                        alt=""
                        src="/logo.svg"
                        width="30"
                        height="30"
                        className="navbar-logo d-inline-block align-top"
                        />{' '}
                        Data Forecast App
                    </Navbar.Brand>
                    <Nav className="me-auto">
                        <Nav.Link href="/home">Home</Nav.Link>
                        <Nav.Link href="/portfolio">Portfolio</Nav.Link>
                        <Nav.Link href="/targets">Targets</Nav.Link>
                    </Nav>
                </Container>
            </Navbar>
            Home Page
            <button onClick={goHome}> Go Home </button>
        </div>
    )
}
export default Home;