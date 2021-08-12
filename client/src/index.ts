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
        location
        year
        wins
        losses
        ties
        other
        games
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
        location
        year
        wins
        losses
        ties
        other
        games
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
        location
        year
      }
    }
  `;

  // some print helpers
  const pad = (s: string): string => {
    const lim = 15 - s.length;
    for (let i = 0; i < lim; i++) {
      s = s + " ";
    }
    return s;
  };

  const printTeam = (t: any) => {
    console.log(`${pad(t.name)} : ${t.year} : ${t.location}`);
  };

  request(serverIP, query, null, requestHeaders)
    // just print the length. should be 150
    .then((data) => {
      let t = data.teamsAll;

      // sort by year then by name so names are grouped together
      // then listed in numeric order
      t = t.sort((a: any, b: any) => {
        return a.year > b.year ? 1 : a.year < b.year ? -1 : 0;
      });

      t = t.sort((a: any, b: any) => {
        return a.name > b.name ? 1 : a.name < b.name ? -1 : 0;
      });

      t.forEach((v: any) => {
        printTeam(v);
      });
    })
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
