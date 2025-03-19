import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  stages: [
    { duration: "10s", target: 1000 },
    { duration: "30s", target: 2083 },
    { duration: "10s", target: 0 },
  ],
};

const BASE_URL = "http://localhost:8080";

export default function () {
  let res = http.get(`${BASE_URL}/mail-address`);
  check(res, {
    "GET /mail-address status es 200": (r) => r.status === 200,
  });

  let createPayload = JSON.stringify({
    address: `test${Math.random()}@example.com`,
    active: Math.random() > 0.5,
  });
  let createHeaders = { "Content-Type": "application/json" };
  let createRes = http.post(`${BASE_URL}/mail-address`, createPayload, {
    headers: createHeaders,
  });
  check(createRes, {
    "POST /mail-address status es 201": (r) => r.status === 201,
  });

  let mailId = createRes.json("id");

  let updatePayload = JSON.stringify({
    address: `test${Math.random()}@example.com`,
    active: Math.random() > 0.5,
  });
  let updateRes = http.put(
    `${BASE_URL}/mail-address/${mailId}`,
    updatePayload,
    { headers: createHeaders },
  );
  check(updateRes, {
    "PUT /mail-address/{id} status es 200": (r) => r.status === 200,
  });

  // let deleteRes = http.del(`${BASE_URL}/mail-address/${mailId}`);
  // check(deleteRes, {
  //   "DELETE /mail-address/{id} status es 204": (r) => r.status === 204,
  // });

  sleep(1);
}
