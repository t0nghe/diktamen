function tokenExists() {
  const token = window.localStorage.getItem("token");
  if (!token || token === "") {
    return false;
  } else {
    return true;
  }
}

export default tokenExists;
