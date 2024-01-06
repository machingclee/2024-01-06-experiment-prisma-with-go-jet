-- CreateTable
CREATE TABLE "student" (
    "id" SERIAL NOT NULL,
    "firstName" TEXT NOT NULL,
    "lastName" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "active" INTEGER NOT NULL,
    "CreatedAt" DOUBLE PRECISION NOT NULL,
    "UpdatedAt" DOUBLE PRECISION NOT NULL,

    CONSTRAINT "student_pkey" PRIMARY KEY ("id")
);
