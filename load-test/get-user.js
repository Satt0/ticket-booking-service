import http from "k6/http";
import { sleep, check } from "k6";

export let options = {
  vus: 2000, // Number of virtual users
  duration: "10s", // Duration of the test
};

export default function () {
  // Replace with your actual API URL
  const id = (Math.floor(Date.now() * Math.random()) % 4) + 1;
//   const id =2
  const url = "http://localhost:8000/users/by-id?id=" + id;
  // Send GET request
  let { status } = http.get(url);

  // Check the response status
  if (status !== 200) {
    console.log(status, id);
  }
  check(status, {
    "is status 200": (status) => status === 200,
  });

  // Optional: Sleep to simulate user think time
  sleep(1);
}
