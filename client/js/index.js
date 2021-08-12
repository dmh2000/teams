"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.requestHeaders = exports.serverIP = void 0;
const graphql_request_1 = require("graphql-request");
// server ip, port and url
exports.serverIP = "http://172.17.59.222:8080/query";
// origin for CORS
exports.requestHeaders = {
    Origin: exports.serverIP,
};
// query the server for a team[] given a name
function queryByName(name) {
    return __awaiter(this, void 0, void 0, function* () {
        const query = graphql_request_1.gql `
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
        graphql_request_1.request(exports.serverIP, query, null, exports.requestHeaders)
            .then((data) => console.log(data))
            .catch((err) => {
            console.log(err);
        });
    });
}
// query the server for a team[] by ID
function queryByID(id) {
    return __awaiter(this, void 0, void 0, function* () {
        const query = graphql_request_1.gql `
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
        graphql_request_1.request(exports.serverIP, query, null, exports.requestHeaders)
            .then((data) => console.log(data))
            .catch((err) => {
            console.log(err);
        });
    });
}
// query the server for a team[] by ID
function queryAll() {
    return __awaiter(this, void 0, void 0, function* () {
        const query = graphql_request_1.gql `
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
        graphql_request_1.request(exports.serverIP, query, null, exports.requestHeaders)
            // just print the length. should be 150
            .then((data) => console.log("teamsAll : ", data.teamsAll.length))
            .catch((err) => {
            console.log(err);
        });
    });
}
function main() {
    return __awaiter(this, void 0, void 0, function* () {
        yield queryByName("Nationals");
        yield queryByID("WAS");
        yield queryAll();
    });
}
main();
