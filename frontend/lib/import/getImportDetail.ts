import { apiKey, endPoint } from "@/constants";
import useSWR from "swr";

export default function getImportNoteDetail({
  idNote,
  token,
}: {
  idNote: string;
  token: string;
}) {
  const fetcher = (url: string) =>
    fetch(url, {
      headers: {
        accept: "application/json",
        Authorization: `Bearer ${token}`,
      },
      cache: "no-store",
    })
      .then((res) => {
        return res.json();
      })
      .then((json) => json.data);
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/importNotes/${idNote}`,
    fetcher
  );

  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
