import {
  Category,
  Customer,
  ExportNote,
  ImportNote,
  ImportNoteDetail,
  Ingredient,
  IngredientDetail,
  IngredientForChoose,
  MeasureUnit,
  Product,
  Role,
  RoleFunction,
  SidebarItem,
  Staff,
  StatusNote,
} from "@/types";
import { LuHome } from "react-icons/lu";
import { MdOutlineWarehouse } from "react-icons/md";
import { GoPeople, GoPerson } from "react-icons/go";
import { z } from "zod";

export const apiKey =
  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOiJnM1cyMUE3U1IiLCJyb2xlIjoiIn0sImV4cCI6MTcwMzU1NTY5OCwiaWF0IjoxNzAzNDY5Mjk4fQ.zm-7b5WY4b98_RUuwy-9HSyYNMAzqtOnkw-Z0aOwPSI";
export const endPoint = "http://localhost:8080/v1";

export const required = z.string().min(1, "Không để trống trường này");

export const products: Product[] = [
  {
    id: "ABC123",
    name: "Tra sua",
    price: 15000,
    status: "active",
    idCate: "14",
    image:
      "https://img.freepik.com/free-photo/green-tea-iced-tall-glass-with-cream-topped-with-iced-green-tea-decorated-with-green-tea-powder_1150-22922.jpg?w=740&t=st=1700729728~exp=1700730328~hmac=9a8b1296341e2c7136d1468e7f0c9b989c7621ca00876b4335a0c985b315e71e",
  },
  {
    id: "AEF455",
    name: "Ca phe",
    price: 17000,
    status: "active",
    idCate: "12",
    image:
      "https://img.freepik.com/free-photo/close-up-glass-coffee-milk_23-2148254986.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "HI234",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    idCate: "42",
    image:
      "https://img.freepik.com/free-photo/club-sandwich-panini-with-ham-cheese-tomato-herbs_2829-19928.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "SW235",
    name: "Sandwich ca ngu",
    price: 68000,
    status: "active",
    idCate: "42",
    image:
      "https://img.freepik.com/premium-photo/tuna-sandwiches-with-lettuce-tomatoes-pickles-onions-slate-tray_483766-284.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "GHJ123",
    name: "Sua tuoi",
    price: 12000,
    status: "active",
    idCate: "12",
    image:
      "https://img.freepik.com/free-photo/milk-pourred-glass-with-cookies_23-2148356806.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=ais",
  },
  {
    id: "AFH123",
    name: "Sua chua",
    price: 18000,
    status: "active",
    idCate: "18",
    image:
      "https://img.freepik.com/free-photo/strawberries-with-whipped-cream-bowl-dark-background_1142-50619.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "AEF355",
    name: "Ca phe",
    price: 17000,
    status: "active",
    idCate: "12",
    image:
      "https://img.freepik.com/free-photo/cold-coffee-drink_144627-18369.jpg?w=740&t=st=1697954121~exp=1697954721~hmac=5b1b188e7f2cb863d08f826a31f99074fe923b6371119e1c652037a1458ef27d",
  },
  {
    id: "ABC243",
    name: "Sandwich",
    price: 32000,
    status: "inactive",
    idCate: "42",
    image:
      "https://img.freepik.com/free-photo/top-view-pita-with-avocado-fried-egg-plate_23-2148749157.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "SINHTO1",
    name: "Sinh to dau",
    price: 37000,
    status: "active",
    idCate: "23",
    image:
      "https://img.freepik.com/premium-photo/strawberry-smoothie-cup-isolated-white_79161-547.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "SINHTO2",
    name: "Sinh to chuoi",
    price: 35000,
    status: "active",
    idCate: "23",
    image:
      "https://img.freepik.com/free-photo/delicious-banana-milkshake_144627-5649.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=ais",
  },
  {
    id: "SINHTO3",
    name: "Sinh to xoai",
    price: 32000,
    status: "active",
    idCate: "23",
    image:
      "https://img.freepik.com/premium-photo/glass-tasty-mango-smoothie-tablemango-smoothie-with-mango-fresh-mango-iceglass-tasty-m_1016228-5865.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "TEA112",
    name: "Matcha Latte",
    price: 35000,
    status: "active",
    idCate: "12",
    image:
      "https://img.freepik.com/premium-photo/glass-green-tea-matcha-latte-green-background_547296-3446.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "JUI123",
    name: "Nuoc chanh",
    price: 35000,
    status: "active",
    idCate: "25",
    image:
      "https://img.freepik.com/free-photo/lemonade-with-lemon-mint-glass-jar-gray-background_1142-50651.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "JUI124",
    name: "Nuoc ep cu cai do",
    price: 35000,
    status: "active",
    idCate: "25",
    image:
      "https://img.freepik.com/free-photo/wooden-tray-with-beetroot-juice_23-2148306956.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
  {
    id: "JUI125",
    name: "Nuoc ep ca rot",
    price: 3000,
    status: "active",
    idCate: "25",
    image:
      "https://img.freepik.com/free-photo/front-view-fresh-organic-carrot-juice_23-2148306958.jpg?size=626&ext=jpg&ga=GA1.1.1850966582.1695100738&semt=sph",
  },
];

export const categories: Category[] = [
  {
    id: "12",
    name: "Ca phe",
    quantity: 5,
  },
  {
    id: "14",
    name: "Tra sua",
    quantity: 5,
  },
  {
    id: "18",
    name: "Sua chua",
    quantity: 5,
  },
  {
    id: "23",
    name: "Sinh to",
    quantity: 6,
  },
  {
    id: "25",
    name: "Nuoc ep",
    quantity: 1,
  },
  {
    id: "42",
    name: "Sandwich",
    quantity: 14,
  },
];

export const measureUnits: MeasureUnit[] = [
  {
    id: "1",
    name: "g",
  },
  {
    id: "2",
    name: "ml",
  },
  {
    id: "3",
    name: "kg",
  },
  {
    id: "4",
    name: "l",
  },
  {
    id: "5",
    name: "đơn vị",
  },
];

export const ingredientForChoose: IngredientForChoose[] = [
  {
    id: "1",
    name: "Sua",
    unitId: "2",
  },
  {
    id: "2",
    name: "Duong",
    unitId: "1",
  },
];

export const ingredientDetails: IngredientDetail[] = [
  {
    idIngre: "1",
    quantity: 17,
    expirationDate: new Date(2024, 8, 29),
  },
  {
    idIngre: "1",
    quantity: 3,
    expirationDate: new Date(2023, 12, 29),
  },
];

export const importNotes: ImportNote[] = [
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY2",
    supplierId: "DT01",
    totalPrice: 3720000,
    status: StatusNote.Done,
    createdAt: new Date(2023, 9, 8),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY3",
    supplierId: "DT01",
    totalPrice: 4660000,
    status: StatusNote.Cancel,
    createdAt: new Date(2023, 10, 1),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createdAt: new Date(),
    createdBy: {
      id: "NV002",
      name: "Nguyễn Thị Huệ",
    },
    supplier: {
      id: "NV002",
      name: "MilkFarm HT",
      phone: "0987654321",
    },
  },
];

export const statuses = [
  {
    isActive: true,
    label: "Đang giao dịch",
  },
  {
    isActive: false,
    label: "Ngừng giao dịch",
  },
];
export const sidebarItems: SidebarItem[] = [
  {
    title: "Quản lý kho",
    href: "/stock-manage",
    icon: MdOutlineWarehouse,
    submenu: true,
    subMenuItems: [{ title: "Nhập kho", href: "/stock-manage/import" }],
  },
];
