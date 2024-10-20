import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <body>
  <div className="grid">
    <div className="box">
      <Link href={"/books"}>Books List</Link>
    </div>
    {/* <div className="box ">
      <button className="button" onClick="onPressParseSkoobHtml()">Parse Skoob HTML</button>
    </div> */}
  </div>
</body>
  );
}
