import { apiKey, endPoint } from "@/constants";

export default async function getAllImportNoteForExcel({
  limit,
  page,
  token,
}: {
  limit?: number;
  page: string;
  token: string;
}) {
  const url = `${endPoint}/importNotes?page=${page}&limit=${limit ?? "10"}`;
  console.log(url);

  const res = await fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  if (!res.ok) {
    // throw new Error("Failed to fetch data");
    console.error(res);
    return res.json();
  }
  return res.json().then((json) => {
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
