import http from "k6/http";
import { sleep, check } from "k6";

export let options = {
  vus: 2000, // Number of virtual users
  duration: "10s", // Duration of the test
};

export default function () {
  // Replace with your actual API URL
  //   const id =2
  const randomIndex = Math.floor(Math.z() * 4);
  const randomCart = [
    {
      optionId: randomIndex + 1,
      amount: 1,
    },
    // {
    //   optionId: ((randomIndex + 1) % 4) + 1,
    //   amount: 1,
    // },
  ];
  const url = "http://localhost:8008/orders";
  // Send GET request
  let { status, body } = http.post(
    url,
    JSON.stringify({
      eventId: 1,
      cart: randomCart.sort((a, b) => a.optionId - b.optionId),
    })
  );

  // Check the response status
  if (status !== 200) {
    console.log(status, body, randomCart);
  }
  check(status, {
    "is status 200": (status) => status === 200,
  });

  // Optional: Sleep to simulate user think time
  sleep(1);
}
