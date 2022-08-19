import { apolloClient } from "../main";
import gql from "graphql-tag";

export async function checkLoggedIn(): Promise<boolean> {
  // if no token, you are not logged in
  const token = window.localStorage.getItem("token");
  if (!token) {
    return false;
  }

  // with a token, we check the validity of it
  try {
    const res = await apolloClient.query({
      query: gql`
        query isLoggedIn {
          getUsername {
            username
            loggedIn
          }
        }
      `,
    });
    if (
      res &&
      res.data &&
      res.data.getUsername &&
      res.data.getUsername.loggedIn
    ) {
      return true;
    } else {
      // if this token is invalid, remove it
      // so next time we won't need to run this query
      window.localStorage.removeItem("token");
      return false;
    }
  } catch (err) {
    console.log(err);
    return false;
  }
}
