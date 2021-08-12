import { request, gql } from "graphql-request";

// server ip, port and url
export const serverIP = "http://172.17.59.222:8080/query";

// origin for CORS
export const requestHeaders = {
  Origin: serverIP,
};

// query the server for a team[] given a name
async function queryByName(name: string) {
  const query = gql`
    {
      teamsByName(name: "${name}") {
        id
        name
        wins
        losses
        ties
        other
        uuid
      }
    }
  `;

  request(serverIP, query, null, requestHeaders)
    .then((data) => console.log(data))
    .catch((err: any) => {
      console.log(err);
    });
}

// query the server for a team[] by ID
async function queryByID(id: string) {
  const query = gql`
    {
      teamsByID(id: "${id}") {
        id
        name
        wins
        losses
        ties
        other
        uuid
      }
    }
  `;

  request(serverIP, query, null, requestHeaders)
    .then((data) => console.log(data))
    .catch((err: any) => {
      console.log(err);
    });
}

// query the server for a team[] by ID
async function queryAll() {
  const query = gql`
    {
      teamsAll() {
        id
        name
        wins
        losses
        ties
        other
        uuid
      }
    }
  `;

  request(serverIP, query, null, requestHeaders)
    // just print the length. should be 150
    .then((data) => console.log("teamsAll : ", data.teamsAll.length))
    .catch((err: any) => {
      console.log(err);
    });
}

async function main(): Promise<void> {
  await queryByName("Nationals");
  await queryByID("WAS");
  await queryAll();
}

main();
