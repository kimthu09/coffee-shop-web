export type Product = {
  id: string;
  name: string;
  price: number;
  status: "active" | "inactive";
  image?: string;
};

export const products: Product[] = [
  {
    id: "ABC123",
    name: "Tra sua",
    price: 15000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AEF355",
    name: "Ca phe",
    price: 17000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "ABC243",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "GHJ123",
    name: "Sua tuoi",
    price: 12000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "AFH123",
    name: "Sua chua",
    price: 18000,
    status: "active",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
];
