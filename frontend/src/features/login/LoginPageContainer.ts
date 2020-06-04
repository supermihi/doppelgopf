import { connect } from "react-redux";
import LoginPage from "./LoginPage";
import { selectLogin, register } from "app/core/login";

const mapState = selectLogin;
const mapDispatch = {
  signup: register,
};
export default connect(mapState, mapDispatch)(LoginPage);
