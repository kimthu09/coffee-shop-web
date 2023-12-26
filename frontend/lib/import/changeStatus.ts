import { apiKey, endPoint } from "@/constants";
import { StatusNote } from "@/types";
import axios from "axios";

export default async function updateStatus({
  idNote,
  status,
  token,
}: {
  idNote: string;
  status: StatusNote;
  token: string;
}) {
  const url = `${endPoint}/importNotes/${idNote}`;
  const data = {
    status: status,
  };
  console.log(data);
  const headers = {
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    accept: "application/json",
  };

  const res = axios
    .post(url, data, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
