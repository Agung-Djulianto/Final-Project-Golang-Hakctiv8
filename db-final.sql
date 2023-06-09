PGDMP                         {            final    15.2    15.2                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16662    final    DATABASE     |   CREATE DATABASE final WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE final;
                postgres    false            �            1259    16686    comments    TABLE     �   CREATE TABLE public.comments (
    id text NOT NULL,
    user_id text,
    photo_id text,
    massage text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.comments;
       public         heap    postgres    false            �            1259    16674    photos    TABLE        CREATE TABLE public.photos (
    id text NOT NULL,
    title character varying(100) NOT NULL,
    caption character varying(255) NOT NULL,
    photo_url character varying(255) NOT NULL,
    user_id text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.photos;
       public         heap    postgres    false            �            1259    16698    social_media    TABLE     )  CREATE TABLE public.social_media (
    id text NOT NULL,
    name character varying(255) NOT NULL,
    social_media_url character varying(255) NOT NULL,
    user_id text,
    created_at timestamp with time zone,
    updated_ats timestamp with time zone,
    updated_at timestamp with time zone
);
     DROP TABLE public.social_media;
       public         heap    postgres    false            �            1259    16663    users    TABLE     (  CREATE TABLE public.users (
    id text NOT NULL,
    user_name character varying(30) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    age smallint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.users;
       public         heap    postgres    false                      0    16686    comments 
   TABLE DATA           Z   COPY public.comments (id, user_id, photo_id, massage, created_at, updated_at) FROM stdin;
    public          postgres    false    216   �                 0    16674    photos 
   TABLE DATA           `   COPY public.photos (id, title, caption, photo_url, user_id, created_at, updated_at) FROM stdin;
    public          postgres    false    215   C                 0    16698    social_media 
   TABLE DATA           p   COPY public.social_media (id, name, social_media_url, user_id, created_at, updated_ats, updated_at) FROM stdin;
    public          postgres    false    217   K                 0    16663    users 
   TABLE DATA           \   COPY public.users (id, user_name, email, password, age, created_at, updated_at) FROM stdin;
    public          postgres    false    214   �       y           2606    16692    comments comments_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.comments DROP CONSTRAINT comments_pkey;
       public            postgres    false    216            w           2606    16680    photos photos_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.photos DROP CONSTRAINT photos_pkey;
       public            postgres    false    215            {           2606    16704    social_media social_media_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.social_media
    ADD CONSTRAINT social_media_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.social_media DROP CONSTRAINT social_media_pkey;
       public            postgres    false    217            q           2606    16673    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public            postgres    false    214            s           2606    16669    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    214            u           2606    16671    users users_user_name_key 
   CONSTRAINT     Y   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_name_key UNIQUE (user_name);
 C   ALTER TABLE ONLY public.users DROP CONSTRAINT users_user_name_key;
       public            postgres    false    214            }           2606    16710    comments fk_photos_comment    FK CONSTRAINT     {   ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_photos_comment FOREIGN KEY (photo_id) REFERENCES public.photos(id);
 D   ALTER TABLE ONLY public.comments DROP CONSTRAINT fk_photos_comment;
       public          postgres    false    3191    215    216            ~           2606    16693    comments fk_users_comment    FK CONSTRAINT     x   ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_users_comment FOREIGN KEY (user_id) REFERENCES public.users(id);
 C   ALTER TABLE ONLY public.comments DROP CONSTRAINT fk_users_comment;
       public          postgres    false    214    3187    216            |           2606    16681    photos fk_users_photos    FK CONSTRAINT     u   ALTER TABLE ONLY public.photos
    ADD CONSTRAINT fk_users_photos FOREIGN KEY (user_id) REFERENCES public.users(id);
 @   ALTER TABLE ONLY public.photos DROP CONSTRAINT fk_users_photos;
       public          postgres    false    3187    214    215                       2606    16705 "   social_media fk_users_social_media    FK CONSTRAINT     �   ALTER TABLE ONLY public.social_media
    ADD CONSTRAINT fk_users_social_media FOREIGN KEY (user_id) REFERENCES public.users(id);
 L   ALTER TABLE ONLY public.social_media DROP CONSTRAINT fk_users_social_media;
       public          postgres    false    217    214    3187               �   x�m�=�0@�99E;##�%vz��u�	1��);z��dr<v�6(]XAQ��8"(��v4h�Z�e�[�^f�H�TE��h/���b0w�v�U<��k�ޟ��s�˺���Q��BmS���b��7��Z��]����.Q         �   x���=n�0�ڜ�>2;���*i<�0֒l�ۇ��M��OOz�|�U�x�t4\{���x�W��*�r,�V'��5��R�8L5���yl\׷��r�}��D{S���ι���0�j
���ǵ�GP����h2	���5�|���p�U����	��)҉#�mTN�$�ʆe�Jyb0>9�c(߶Lۜq;);����t㑇c8>�&k�XP�O�j�hA7N{p�A~m����L�w         �   x�}���0 �s2w��q�e&��NQ���?L��/��X�w��W(�R�Т

e3���1�쏝|��c>�]��I��V�&�
0����CWCH+��N�/�~�5����l���n���I[+�         �  x�}�Ms�0���+r�֑�d}�T׆B�B�	Ӌ$��`��e�=���mg��>��JE�r� ��-� Fp)�%�K娬��;�8T���b��[��;��U��bv�<�y�����۸p�����E����A�}�P�}�{|������!�$ �;$J�{Dr&З���V�0+��0`H@xN�̡����Z*��c�̾]�N�JU����M�ܞڤQ׆�e:Km=a��v2����ToŏUTD�զY����_�B��<)!a�յ��#4g
�b �5��
D!d:��tSW�ѫ6�o�U�p?�7�Ӭ��Y2G����)��ް?-�'��Q�)����NG��&u<���%Òb�����X�I 1�0���~��aEsk��WA�����[��*���z2~|�����C�M���$����I_5���w{�4�� �RH���?��O�u���H�     