"use client";

import { useState } from "react";
import axios from "axios";

import Image from "next/image";
import { useRouter } from "next/navigation";
import { useToast } from "@/components/ui/use-toast";
import Cookies from "js-cookie";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { endPoint } from "@/constants";
import { useAuth } from "@/hooks/auth-context";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();
  const { toast } = useToast();
  const { login } = useAuth();

  const handleLogin = async (event: any) => {
    event.preventDefault();
    const response = await axios
      .post(endPoint + "/login", {
        email: email,
        password: password,
      })
      .then((res) => {
        const { access_token, refresh_token } = res.data.data;

        const expirationDate = new Date(
          new Date(access_token.created).getTime() + access_token.expiry * 1000
        );
        // Save tokens in cookies
        Cookies.set("accessToken", access_token.token, {
          expires: expirationDate,
        });
        Cookies.set("refreshToken", refresh_token.token, {
          expires: refresh_token.expiry,
        });
        login("");

        // Redirect to the main page
        router.push("/");
        router.refresh();
      })
      .catch((e) => {
        console.log("Login failed:", e.message);
        toast({
          title: "Đăng nhập thất bại",
          description: "Vui lòng kiểm tra lại email và mật khẩu của bạn",
        });
      });
  };

  return (
    <div className="container relative flex pt-20 flex-col items-center justify-center lg:px-0">
      <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[500px]">
        <div className="shadow-lg p-5 rounded-lg border-t-4 border-orange-400">
          <div className="flex flex-col items-center space-y-2 text-center pb-5">
            <Image
              src="/android-chrome-192x192.png"
              alt="logo"
              width={50}
              height={50}
            ></Image>
            <h1 className="text-xl font-bold my-4">Đăng nhập vào tài khoản</h1>
          </div>
          <form onSubmit={handleLogin} className="flex flex-col gap-3">
            <div className="pb-1">
              <Input
                onChange={(e) => setEmail(e.target.value)}
                type="email"
                placeholder="Email"
              />
            </div>
            <div className="pb-5">
              <Input
                onChange={(e) => setPassword(e.target.value)}
                type="password"
                placeholder="Mật khẩu"
              />
            </div>

            <Button className="bg-orange-400 text-white font-bold cursor-pointer px-6 py-2">
              Login
            </Button>
          </form>
        </div>
      </div>
    </div>
  );
};
// export default withLayout(Login, { isLoginPage: true });
export default Login;
