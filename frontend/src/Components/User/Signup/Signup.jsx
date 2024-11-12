import "./Signup.css"

const Signup = () => {
    return (
        <div id="signup-form-wrapper">
            <form action="http://localhost:8080/adduser" method="post">
            {/* <form action="https://httpdump.app/dumps/5f468417-215f-4f39-8548-1c4245e5b7a9" method="post"> */}
                <label>Name</label>
                <input type="text" name="username" id="" />
                <label htmlFor="">Password</label>
                <input type="password" name="password" id="" />
                <input type="submit" value="Envoyer" id="signup-form-submit-btn" />
            </form>
        </div>
    )
}

export default Signup