"use client";
import { createContext, useContext, useState, useEffect } from "react";
import Cookies from "js-cookie";

interface User {
  // Define your user properties here
  username: string;
  // ...
}

interface AuthContextProps {
  user: User | null;
  login: (token: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<User | null>(null);

  const login = (token: string) => {
    // Set the token in cookies
    // Cookies.set("token", token, { expires: 7, path: "/" });

    // For simplicity, you can set a default user here
    setUser({ username: "example" });
  };

  const logout = () => {
    // Remove the token from cookies
    Cookies.remove("accessToken", { path: "/" });
    setUser(null);
    window.location.reload();
  };

  useEffect(() => {
    // Check for existing token when the component mounts
    const token = Cookies.get("token");
    if (token) {
      // For simplicity, you can set a default user here
      setUser({ username: "example" });
    }
  }, []);

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextProps => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
